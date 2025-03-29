package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// Post 文章或页面
type Post struct {
	Id         uint64     `json:"id"`                     // 数据ID
	Title      string     `json:"title"`                  // 标题
	Url        string     `json:"url"`                    // 文章地址
	Summary    string     `json:"summary"`                // 摘要
	SourceId   uint64     `json:"sourceId"`               // 主图资源ID
	PostType   string     `json:"postType"`               // 文章类型，枚举：0_页面 1_文章
	CategoryId uint64     `json:"categoryId"`             // 文章分组
	PushAt     *time.Time `json:"pushAt"`                 // 发布时间，可以自动任务抓取，不配代表不自动发布
	Status     string     `json:"status"`                 // 状态，枚举：0_草稿 1_发布 2_锁定
	CreateAt   *time.Time `json:"createAt"`               // 创建时间
	UpdateAt   *time.Time `json:"updateAt"`               // 更新时间
	Views      uint64     `json:"views" gorm:"<-:create"` // 总浏览量 / 允许读
	Goods      uint64     `json:"goods" gorm:"<-:create"` // 总支持量 / 允许读
	Hots       uint64     `json:"hots" gorm:"<-:create"`  // 近期热度，浏览量+支持*100，每周零点/100 / 允许读
}

// PostTable 文章或页面泛型造器
var PostTable actuator.Table[Post]

// DataBase 实现指定数据库
func (t Post) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t Post) TableName() string {
	return "post"
}

// ToCategory 迁移到新分组
func (t Post) ToCategory(id, toId uint64) (err error) {
	now := time.Now()
	r := blogDb.Table(t.TableName()).Where("category_id = ?", id).Updates(map[string]any{
		"category_id": toId,
		"update_at":   &now,
	})
	err = r.Error
	return
}

// GetPosts 迁移到新分组
func (t Post) GetPosts() (res []*Post, err error) {
	r := blogDb.Table(t.TableName()).Where("status = ?", "1").Where("post_type", "1").Order("push_at DESC").Find(&res)
	err = r.Error
	return
}

// PostUpdateData 文章或页面
type PostUpdateData struct {
	Id    uint64 // 数据ID
	Views uint64 // 总浏览量 / 允许读
	Goods uint64 // 总支持量 / 允许读
	Hots  uint64 // 近期热度，浏览量+支持*100，每周零点/100 / 允许读
}

// UpdatePostDataBatch 批量更新文章数据
func (t Post) UpdatePostDataBatch(req []*PostUpdateData) (err error) {
	db := t.DataBase()
	// 启用事务
	return db.Transaction(func(tx *gorm.DB) error {
		for _, item := range req {
			// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
			if err := tx.Table(t.TableName()).Where("id = ?", item.Id).Update("views", item.Views).Update("goods", item.Goods).Update("hots", item.Hots).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}
		}
		// 返回 nil 提交事务
		return nil
	})
}
