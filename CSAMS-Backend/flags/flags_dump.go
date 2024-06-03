package flags

import (
	"CSAMS-Backend/global"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// Dump 导出mysql数据
func Dump() {
	mysql := global.Config.Mysql
	//获取当前时间
	timer := time.Now().Format("20060102")
	//拼装文件名
	sqlPath := fmt.Sprintf("%s_%s.sql", mysql.DB, timer)
	// 构建 mysqldump 命令
	cmd := exec.Command("mysqldump", "-u"+mysql.User, "-p"+mysql.Password, mysql.DB)
	// 创建导出文件
	outFile, err := os.Create(sqlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// 将 mysqldump 的输出写入到文件中
	cmd.Stdout = outFile
	// 执行命令
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sql文件 %s 导出成功", sqlPath)
}
