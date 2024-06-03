package flags

import (
	"CSAMS-Backend/global"
	"log"
	"os"
	"strings"
)

// Load 导入mysql
func Load(sqlPath string) {
	//读取文件
	byteData, err := os.ReadFile(sqlPath)
	if err != nil {
		log.Fatalf("%s err: %s", sqlPath, err.Error())
	}
	//分割数据 一定要按照\r\n分割
	sqlList := strings.Split(string(byteData), ";\r\n")
	for _, sql := range sqlList {
		//去除字符串开头和结尾的空白符
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		//执行sql语句
		err = global.DB.Exec(sql).Error
		if err != nil {
			log.Printf("%s err:%s", sql, err.Error())
			continue
		}
	}
	log.Printf("%s sql导入成功", sqlPath)
}
