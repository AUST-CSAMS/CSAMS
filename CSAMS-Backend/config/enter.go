package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	System System `yaml:"system"`
	Upload Upload `yaml:"upload"`
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
}
