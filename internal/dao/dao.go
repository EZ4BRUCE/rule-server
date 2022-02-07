package dao

import "gorm.io/gorm"

// Dao封装一个db引擎实例用于执行model方法
type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
