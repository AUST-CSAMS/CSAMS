package flags

import "flag"

type Option struct {
	DB     bool   // 初始化数据库 -db
	Import string // 导入数据库文件 -import
	Export bool   // 导出数据库 -export
}

// Parse 解析命令行参数
func Parse() (option *Option) {
	option = new(Option)
	//解析命令行字段并保存到option中
	flag.BoolVar(&option.DB, "db", false, "初始化数据库")
	flag.BoolVar(&option.Export, "export", false, "导出sql数据库")
	flag.StringVar(&option.Import, "import", "", "导入sql数据库")
	flag.Parse()
	return option
}

// Run 根据命令执行不同的函数
func (option Option) Run() bool {
	if option.DB {
		DB()
		return true
	}
	if option.Import != "" {
		Load(option.Import)
		return true
	}
	if option.Export {
		Dump()
		return true
	}
	return false
}
