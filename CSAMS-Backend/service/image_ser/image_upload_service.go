package image_ser

import (
	"CSAMS-Backend/global"
	"CSAMS-Backend/utils"
	"fmt"
	"mime/multipart"
	"path"
	"strings"
)

var WhiteImageList = []string{
	"jpg",
	"png",
	"jpeg",
	"ico",
	"tiff",
	"gif",
	"svg",
	"webp",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// ImageUploadService 图片上传
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	//拼装文件路径
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	res.FileName = filePath
	// 文件白名单判断
	nameList := strings.Split(fileName, ".")
	//转小写
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return
	}
	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%dMB ", size, global.Config.Upload.Size)
		return
	}
	res.Msg = "图片上传成功"
	res.IsSuccess = true
	filePath = "/" + filePath

	return
}
