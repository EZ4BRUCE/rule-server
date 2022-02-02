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

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
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

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	// 用语句建表
	global.DBEngine.Exec("CREATE TABLE IF NOT EXISTS `function` ( `id` int(10) unsigned NOT NULL AUTO_INCREMENT, `type` varchar(255) NOT NULL COMMENT '类型', `threshold` double NOT NULL COMMENT '阈值', `description` varchar(255) COMMENT '描述', PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='聚合函数表';")
	global.DBEngine.Exec("CREATE TABLE IF NOT EXISTS `rule` ( `id` int(10) unsigned NOT NULL AUTO_INCREMENT, `level` varchar(255) NOT NULL COMMENT '告警等级', `action` varchar(255) NOT NULL COMMENT '告警行为', `description` varchar(255) DEFAULT '' COMMENT '描述', PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警规则表';")
	global.DBEngine.Exec("CREATE TABLE IF NOT EXISTS `aggregator` ( `id` int(10) unsigned NOT NULL AUTO_INCREMENT, `name` varchar(255) UNIQUE COMMENT '聚合器名称', `metric` varchar(255) DEFAULT '' COMMENT '聚合指标', `function_id` int(10) unsigned COMMENT '关联的聚合函数id', `rule_id` int(10) unsigned  COMMENT '关联的告警规则id', PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='聚合器表';")
	loadDeafult()
	return nil
}

// 载入默认配置
func loadDeafult() {
	// 载入function默认配置
	global.DBEngine.Exec("INSERT INTO `function` VALUES (1, 'MAX', 0.9, '最大值类型函数，聚合数据的最大值不可大于阈值');")
	global.DBEngine.Exec("INSERT INTO `function` VALUES (2, 'AVG', 0.9, '平均值类型函数，聚合数据的平均值值不可大于阈值');")
	global.DBEngine.Exec("INSERT INTO `function` VALUES (3, 'MIN', 0.1, '最小值类型函数，聚合数据的最小值不可小于阈值');")
	global.DBEngine.Exec("INSERT INTO `function` VALUES (4, 'SUM', 4.5, '求和类型函数，聚合结果值不可大于阈值');")
	// 载入rule默认配置
	global.DBEngine.Exec("INSERT INTO `rule` VALUES (1, 'INFO', 'MESSAGE', '特定情况发送短信通知');")
	global.DBEngine.Exec("INSERT INTO `rule` VALUES (2, 'WARNING', 'EMAIL', '特定情况发送邮件到指定邮箱');")
	global.DBEngine.Exec("INSERT INTO `rule` VALUES (3, 'FATAL', 'PHONE', '特定情况打电话');")
	// 载入aggregator默认配置
	global.DBEngine.Exec("INSERT INTO `aggregator` VALUES (11, '内存占用最大值监测', 'memory_used', 1, 2);")
	global.DBEngine.Exec("INSERT INTO `aggregator` VALUES (12, '内存占用平均值监测', 'memory_used', 2, 2);")
	global.DBEngine.Exec("INSERT INTO `aggregator` VALUES (13, '内存占用最小值监测', 'memory_used', 3, 1);")
	global.DBEngine.Exec("INSERT INTO `aggregator` VALUES (14, '内存占用求和监测', 'memory_used', 4, 3);")
	global.DBEngine.Exec("INSERT INTO `aggregator` VALUES (21, 'CPU占用最大值监测', 'cpu_rate', 1, 2);")
	global.DBEngine.Exec("INSERT INTO `aggregator` VALUES (22, 'CPU占用平均值监测', 'cpu_rate', 2, 2);")
	global.DBEngine.Exec("INSERT INTO `aggregator` VALUES (23, 'CPU占用最小值监测', 'cpu_rate', 3, 1);")
	global.DBEngine.Exec("INSERT INTO `aggregator` VALUES (24, 'CPU占用求和监测', 'cpu_rate', 4, 3);")
}
