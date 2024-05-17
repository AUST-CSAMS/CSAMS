package utils

import (
	"CSAMS-Backend/global"
	"log"
)

func PrintSystem() {

	ip := global.Config.System.Host
	port := global.Config.System.Port

	if ip == "0.0.0.0" {
		ipList := GetIPList()
		for _, i := range ipList {
			log.Println("CSAMS-Backend 运行在： http://%s:%d/api", i, port)
			log.Println("CSAMS-Backend api文档 运行在： http://%s:%d/swagger/index.html#", i, port)
		}
	} else {
		log.Println("CSAMS-Backend 运行在： http://%s:%d/api", ip, port)
		log.Println("CSAMS-Backend api文档 运行在： http://%s:%d/swagger/index.html#", ip, port)
	}

}
