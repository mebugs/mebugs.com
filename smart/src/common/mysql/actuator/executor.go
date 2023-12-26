package actuator

import (
	"fmt"

	"gorm.io/gorm"
)

// MYSQL 执行器
// 常见数据库执行API

// Executor 泛型接口
type Executor interface {
	TableName() string
	DataBase() *gorm.DB
}

// Table 泛型执行器
type Table[T Executor] string

// FindAll 查询表所有数据
func (t Table[T]) FindAll() (res []*T, err error) {
	var exe T
	r := exe.DataBase().Find(&res)
	err = r.Error
	return
}

// FindOneById 查询表指定ID的数据
func (t Table[T]) FindOneById(id uint64) (res T, err error) {
	var exe T
	r := exe.DataBase().Where("id = ?", id).First(&res)
	err = r.Error
	return
}

// FindOneByObject 查询表指定对象的数据（传入指针）
func (t Table[T]) FindOneByObject(req any) (res T, err error) {
	var exe T
	r := exe.DataBase().Where(req).First(&res)
	err = r.Error
	return
}

// FindByIds 查询表指定ID的数据集
func (t Table[T]) FindByIds(ids []uint64) (res []*T, err error) {
	var exe T
	r := exe.DataBase().Where("id IN ?", ids).Find(&res)
	err = r.Error
	return
}

// FindByObject 查询表指定ID的数据集
func (t Table[T]) FindByObject(req any) (res []*T, err error) {
	var exe T
	r := exe.DataBase().Where(req).Find(&res)
	err = r.Error
	return
}

// Page 平台的公共分页查询方法
func (t Table[T]) Page(query *Query) (total int64, list []*T, err error) {
	var exe T
	// 提交表名
	query.Table = exe.TableName()
	// 先查询数量
	total, err = query.countByQuery(exe.DataBase())
	if err != nil || total == 0 {
		return
	}
	// 开始查询分页数据
	err = query.findByQuery(exe.DataBase(), &list)
	if err != nil {
		return
	}
	return
}

// InsertOne 插入一条数据
func (t Table[T]) InsertOne(req any) (err error) {
	var exe T
	r := exe.DataBase().Create(req)
	err = r.Error
	return
}

// InsertBatch 插入多条数据
func (t Table[T]) InsertBatch(req any) (err error) {
	var exe T
	r := exe.DataBase().Create(req)
	err = r.Error
	return
}

// UpdateOne 更新一条数据（根据ID）
func (t Table[T]) UpdateOne(req any) (err error) {
	var exe T
	r := exe.DataBase().Save(req)
	err = r.Error
	return
}

// DeleteOne 删除一条数据（根据ID）
func (t Table[T]) DeleteOne(id uint64) (err error) {
	var exe T
	r := exe.DataBase().Exec(fmt.Sprintf("DELETE FROM %s WHERE `id`= ?", exe.TableName()), id)
	err = r.Error
	return
}

// DeleteByIds 删除多条数据（根据ID）
func (t Table[T]) DeleteByIds(ids []uint64) (err error) {
	var exe T
	r := exe.DataBase().Exec(fmt.Sprintf("DELETE FROM %s WHERE `id`IN ?", exe.TableName()), ids)
	err = r.Error
	return
}
