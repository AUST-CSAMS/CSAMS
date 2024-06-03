package image_api

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/models/res"
	"CSAMS-Backend/service/image_ser"
	"CSAMS-Backend/utils"
	"CSAMS-Backend/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

// ImageUploadDataView 上传单个图片，返回图片的url
func (ImageApi) ImageUploadDataView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		log.Print(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//获取token信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	fileName := file.Filename
	basePath := global.Config.Upload.Path
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	//判断是否在白名单内
	if !utils.InList(suffix, image_ser.WhiteImageList) {
		res.FailWithMessage("非法文件", c)
		return
	}
	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		msg := fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%dMB ", size, global.Config.Upload.Size)
		res.FailWithMessage(msg, c)
		return
	}
	//判断文件目录是否存在
	dirList, err := os.ReadDir(basePath)
	if err != nil {
		res.FailWithMessage("文件目录不存在", c)
		return
	}
	// 创建这个文件夹 /uploads/file/ID
	if !isInDirEntry(dirList, claims.ID) {
		err := os.Mkdir(path.Join(basePath, claims.ID), 077)
		if err != nil {
			log.Print(err)
		}
	}
	// 1.如果存在重名，就加随机字符串 时间戳
	// 2.直接+时间戳
	now := time.Now().Format("20060102150405")
	fileName = nameList[0] + "_" + now + "." + suffix
	//拼装文件名
	filePath := path.Join(basePath, claims.ID, fileName)
	//保存到本地
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData("/"+filePath, c)
	return

}

// isInDirEntry 判断文件名是否在目录中
func isInDirEntry(dirList []os.DirEntry, name string) bool {
	for _, entry := range dirList {
		if entry.Name() == name && entry.IsDir() {
			return true
		}
	}
	return false
}
