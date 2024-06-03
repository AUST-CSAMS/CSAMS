package utils

import (
	"log"
	"net"
)

// GetIPList 获取本地计算机上所有网络接口的 IPv4 地址列表
func GetIPList() (ipList []string) {
	//返回本地系统所有网络接口的列表
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i2 := range interfaces {
		//获取该网络接口上的所有地址
		address, err := i2.Addrs()
		if err != nil {
			log.Println(err)
			continue
		}
		for _, addr := range address {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			// IP 地址转换为 IPv4 格式
			ip4 := ipNet.IP.To4()
			if ip4 == nil {
				continue
			}
			ipList = append(ipList, ip4.String())
		}
	}
	return
}
