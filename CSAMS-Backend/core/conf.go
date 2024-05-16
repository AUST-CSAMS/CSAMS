package core

import (
	"CSAMS-Backend/config"
	"CSAMS-Backend/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"log"
	"os"
)

const ConfigFile = "config.yaml"

// InitConf 读取yaml文件的配置
func InitConf() {
	c := &config.Config{}
	//读取文件 读取出来的内容为字节类型的切片
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	//反序列化 把读取出来的内容转换为对应的格式
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}

// SetYaml 保存对yaml文件的修改
func SetYaml() error {
	//序列化 把对应的内容转化为字节类型的切片
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	//写入文件 把转化的内容写入对应文件WriteFile(写入哪个文件，写入的内容，权限)
	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	log.Println("配置文件修改成功")
	return nil
}
