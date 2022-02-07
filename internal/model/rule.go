package model

import "gorm.io/gorm"

// 告警规则Rule对应的数据库字段信息
type Rule struct {
	Id          uint32 `gorm:"primaryKey;column:id" json:"id"`
	Level       string `gorm:"level" json:"level"`
	Action      string `gorm:"action" json:"action"`
	Description string `gorm:"description" json:"description"`
}

// 创建告警规则至数据库
func (r Rule) Create(db *gorm.DB) error {
	return db.Create(&r).Error
}

// 删除数据库中特定告警规则
func (r Rule) Delete(db *gorm.DB) error {
	return db.Where("id = ?", r.Id).Delete(&r).Error
}

// 获取数据库中特定告警规则
func (r Rule) Get(db *gorm.DB) (Rule, error) {
	var Rule Rule
	err := db.Where("id = ?", r.Id).First(&Rule).Error
	return Rule, err

}

// 更新数据库中特定告警规则
func (r Rule) Update(db *gorm.DB) error {
	var temp Rule
	temp.Id = r.Id
	return db.Model(&temp).Updates(r).Error
}

// 获取数据库中所有告警规则
func (r Rule) List(db *gorm.DB) ([]Rule, error) {
	var rules []Rule
	err := db.Find(&rules).Error
	if err != nil {
		return nil, err
	}
	return rules, nil
}
