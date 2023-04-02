package main

import (
	"fmt"
	"my-admin/config"
	"my-admin/global"
	"my-admin/logs"
	"my-admin/model"
	"my-admin/pkg/minio"
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

	global.MinioClient, err = minio.InitMinio(global.Config.Minio.IP, global.Config.Minio.Port, global.Config.Minio.AccessKeyID, global.Config.Minio.SecretAccessKey, false)
	if err != nil {
		global.Logger.Errorf("failed to init minio client, err:%+v\n", err)
	}
	service.RegisterTable(global.DBClient)
	fmt.Println(service.Register(model.User{
		UserName: "admin",
		Phone:    "13888888888",
		Email:    "test@163.com",
		Password: "123456",
		Avatar:   "http://minio.test.com:9000/test1/1677742742506.jpg",
	}))

	err = router.StartServer()
	if err != nil {
		global.Logger.Errorf("webserver failed, err : %+vreturn.\n", err)
	}
	// 定时任务
	// 插件manager
}
