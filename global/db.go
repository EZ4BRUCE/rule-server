package global

import "gorm.io/gorm"

var (
	// 全局数据库操作引擎
	DBEngine *gorm.DB
)
