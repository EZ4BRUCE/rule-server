package main

import (
	"log"
	"net/http"
	"time"

	"github.com/EZ4BRUCE/rule-server/global"
	"github.com/EZ4BRUCE/rule-server/internal/model"
	"github.com/EZ4BRUCE/rule-server/internal/routers"
	"github.com/EZ4BRUCE/rule-server/pkg/setting"
	"github.com/gin-gonic/gin"
)

// 项目配置初始化，仅在程序开始时执行一次
func init() {
	log.SetPrefix("[Rule-Server]")
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

}

// @title 告警规则配置系统
// @version 1.0
// @description Athena告警系统的告警规则配置子系统
// @termsOfService https://github.com/EZ4BRUCE/rule-server
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	// 不使用gin默认的引擎，而是将配置好的引擎作为handler传入http.Server
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// 从configs中载入global配置
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 根据global的设置初始化数据库
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
