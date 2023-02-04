package main

import (
	"flag"
	"go-email/initialize"
	"go-email/pkg/snowflake"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "f", "./config/config-dev.yaml", "配置文件路径")
	flag.Parse()

	// 初始化配置
	initialize.InitConfig(configPath)
	// 初始化日志
	initialize.InitLogger()
	// 初始化雪花算法结点
	snowflake.Init()
	// 初始化数据库
	initialize.InitDB()
	// 初始化缓存
	initialize.InitRedis()
}
