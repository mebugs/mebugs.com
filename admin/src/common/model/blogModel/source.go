package blogModel

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// SourceDoReq 资源配置表 通用请求，创建&编辑可复用的字段
type SourceDoReq struct {
	Name         string   `json:"name" binding:"required" example:"xxx"`         // 图片资源名
	FileStrArray []string `json:"fileStrArray" binding:"required" example:"xxx"` // 图片Base64数据，主图3个 后面均2个
}

// SourceAddReq 资源配置表 创建请求，酌情从通用中摘出部分字段
type SourceAddReq struct {
	FileType string `json:"fileType" binding:"oneof='0' '1' '2'" example:"0"` // 图片类型，枚举：0_主图 1_图标 2_正文
	SourceDoReq
}

// SourceEditReq 资源配置表 编辑请求，酌情从通用中摘出部分字段
type SourceEditReq struct {
	Id       uint64 `json:"id" binding:"required" example:"1"`                // 数据ID
	FileType string `json:"fileType" binding:"oneof='0' '1' '2'" example:"0"` // 图片类型，枚举：0_主图 1_图标 2_正文
	SourceDoReq
}

// ToDbReq 资源配置表 创建转数据库
func (r *SourceAddReq) ToDbReq() *blogDB.Source {
	now := time.Now()
	// 获取当前月份字符串
	return &blogDB.Source{
		Id:       0,
		Name:     r.Name,
		FileType: r.FileType,
		Status:   constant.StatusLock, // 先锁定，传完更新
		CreateAt: &now,
		UpdateAt: &now,
	}
}

// SourceGetRes 资源配置表 详情响应
type SourceGetRes struct {
	Id       uint64     `json:"id" example:"1"`          // 数据ID
	Img      string     `json:"img" example:"1"`         // 图片地址
	Name     string     `json:"name" example:"xxx"`      // 图片资源名
	BackEnd  string     `json:"backEnd" example:"jpg"`   // 图片后缀，jpg png gif
	FileType string     `json:"fileType" example:"0"`    // 图片类型，枚举：0_主图 1_图标 2_正文
	FilePath string     `json:"filePath" example:"demo"` // 文件子目录，年份目录
	Status   string     `json:"status" example:"0"`      // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"`                // 创建时间
	UpdateAt *time.Time `json:"updateAt"`                // 更新时间
}

// SourcePageReq 资源配置表 分页请求，根据实际业务替换分页条件字段
type SourcePageReq struct {
	Name     string `json:"name" example:"xxx"`   // 图片资源名
	FileType string `json:"fileType" example:"0"` // 图片类型，枚举：0_主图 1_图标 2_正文
	baseModel.PageReq
}

// SourcePageRes 资源配置表 分页响应，酌情从详情摘出部分字段
type SourcePageRes struct {
	SourceGetRes
}

// ToSourceGetRes 资源配置表 数据库转为详情响应
func ToSourceGetRes(r *blogDB.Source) *SourceGetRes {

	return &SourceGetRes{
		Id:       r.Id,
		Img:      fmt.Sprintf(constant.SourceFileUrl, r.FilePath, fmt.Sprintf("%d", r.Id), r.BackEnd, r.Version),
		Name:     r.Name,
		BackEnd:  r.BackEnd,
		FileType: r.FileType,
		FilePath: r.FilePath,
		Status:   r.Status,
		CreateAt: r.CreateAt,
		UpdateAt: r.UpdateAt,
	}
}

// ToSourcePageRes 资源配置表 数据库转分页响应
func ToSourcePageRes(list []*blogDB.Source) []*SourcePageRes {
	res := make([]*SourcePageRes, len(list))
	for i, r := range list {
		res[i] = &SourcePageRes{
			SourceGetRes: *ToSourceGetRes(r),
		}
	}
	return res
}
