package main

import (
	"my-admin/config"
	"my-admin/global"
	"my-admin/logs"
	"my-admin/pkg/mysql"
	"my-admin/pkg/redis"
	"my-admin/router"
	"my-admin/service"
)

func main() {
	var configFile = "config/config.yaml"
	var err error

	global.Logger = logs.InitLogger("../running.log", global.LogLevel, logs.DevMode)

	isConfigUpdate := make(chan struct{})
	global.Config, err = config.LoadConfig(configFile, isConfigUpdate)
	if err != nil {
		global.Logger.Errorf("failed to load config from:%s, err:%+v\n", configFile, err)
	}
	go func() {
		for {
			select {
			case <-isConfigUpdate:
				global.Logger.Infof("local config file is updated, start to reload config\n")
				global.Config, err = config.ReloadConfig()
				if err != nil {
					global.Logger.Errorf("failed to load config from:%s, err:%+v\n", configFile, err)
					continue
				}
				global.Logger.Infof("successfully to reload new config\n")
			}
		}
	}()

	global.DBClient, err = mysql.InitMySql(global.Config.Mysql.Url, global.Config.Mysql.User, global.Config.Mysql.Passwd, global.Config.Mysql.DbName, global.Config.Mysql.Port)
	if err != nil {
		global.Logger.Errorf("failed to init mysql client, err:%+v\n", err)
	}

	global.RedisClient, err = redis.InitRedis(global.Config.Redis.Addr, global.Config.Redis.Passwd)
	if err != nil {
		global.Logger.Errorf("failed to init redis client, err:%+v\n", err)
	}
	service.RegisterTable(global.DBClient)

	err = router.StartServer()
	if err != nil {
		global.Logger.Errorf("webserver failed, err : %+vreturn.\n", err)
	}
	// 定时任务
	// 插件manager
}
