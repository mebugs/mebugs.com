package actuator

import "gorm.io/gorm"

// MYSQL 执行器
// 常见数据库执行API

// FindOneByObject 实体类作为查询条件查一个
func FindOneByObject(db *gorm.DB, query interface{}, res interface{}) error {
	r := db.Where(query).Find(res)
	if r.Error != nil {
		return r.Error
	}
	return nil
}
