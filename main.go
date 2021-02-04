package main

import (
	"bluebell/controllers"
	"bluebell/controllers/run"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routes"
	"bluebell/settings"
	"fmt"
	"os"

	"go.uber.org/zap"
)

// Go Web 开发脚手架模板

func main() {
	// 1. 加载配置
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		if err := settings.Init2(fileName); err != nil {
			fmt.Printf("init settings failed, err: %s\n", err)
			return
		}
	} else {
		fmt.Println("need config file  eg.config.yaml")
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err: %s\n", err)
		return
	}
	zap.L().Debug("logger init success....")
	defer zap.L().Sync()

	// 3. 初始化mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %s\n", err)
		return
	}
	zap.L().Debug("mysql init success....")
	defer mysql.Close()

	// 4. 初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err: %s\n", err)
		return
	}
	zap.L().Debug("redis init success....")
	defer redis.Close()

	// 雪花算法，分布式id
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineId); err != nil {
		fmt.Printf("init snowflake failed, err: %s\n", err)
		return
	}

	// InitTrans初始化校验翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator translator failed, err=%v\n", err)
	}

	// 5. 注册路由
	r := routes.SetUp(settings.Conf.Mode)

	// 6. 启动服务（优雅关机）
	run.ForeverElegant(r)
}
