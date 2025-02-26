package blogModel

import (
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// UserDoReq 极简用户信息 通用请求，创建&编辑可复用的字段
type UserDoReq struct {
	ClientId string     `json:"clientId" binding:"max=64" example:"demo"`                // 客户端ID（前端生成，清除会丢失）
	Name     string     `json:"name" binding:"max=32" example:"demo"`                    // 用户名
	Code     string     `json:"code" binding:"max=32" example:"demo"`                    // 用户编码（允许访问用户）默认=ID
	Summary  string     `json:"summary" binding:"max=64" example:"demo"`                 // 简介
	ThirdUrl string     `json:"thirdUrl" binding:"max=255" example:"demo"`               // 三方网站（需要审核）
	WaitUrl  string     `json:"waitUrl" binding:"max=255" example:"demo"`                // 待审核三方URL
	SourceId uint64     `json:"sourceId" example:"0"`                                    // 头像资源ID（仅小程序提供）
	OpenId   string     `json:"openId" binding:"max=64" example:"demo"`                  // 小程序账号绑定登录（允许授权转移）
	Status   string     `json:"status" binding:"required,oneof='0' '1' '2'" example:"0"` // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"`                                                // 创建时间
	UpdateAt *time.Time `json:"updateAt"`                                                // 更新时间
}

// UserAddReq 极简用户信息 创建请求，酌情从通用中摘出部分字段
type UserAddReq struct {
	UserDoReq
}

// UserEditReq 极简用户信息 编辑请求，酌情从通用中摘出部分字段
type UserEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	UserDoReq
}

// ToDbReq 极简用户信息 创建转数据库
func (r *UserAddReq) ToDbReq() *blogDB.User {
	now := time.Now()
	return &blogDB.User{
		Id:       0,
		ClientId: r.ClientId,
		Name:     r.Name,
		Code:     r.Code,
		Summary:  r.Summary,
		ThirdUrl: r.ThirdUrl,
		WaitUrl:  r.WaitUrl,
		SourceId: r.SourceId,
		OpenId:   r.OpenId,
		Status:   r.Status,
		CreateAt: &now,
		UpdateAt: &now,
	}
}

// ToDbReq 极简用户信息 更新转数据库
func (r *UserEditReq) ToDbReq(d *blogDB.User) {
	now := time.Now()
	d.ClientId = r.ClientId
	d.Name = r.Name
	d.Code = r.Code
	d.Summary = r.Summary
	d.ThirdUrl = r.ThirdUrl
	d.WaitUrl = r.WaitUrl
	d.SourceId = r.SourceId
	d.OpenId = r.OpenId
	d.Status = r.Status
	d.CreateAt = &now
	d.UpdateAt = &now
}

// UserGetRes 极简用户信息 详情响应
type UserGetRes struct {
	Id       uint64     `json:"id" example:"1"`          // 数据ID
	ClientId string     `json:"clientId" example:"demo"` // 客户端ID（前端生成，清除会丢失）
	Name     string     `json:"name" example:"demo"`     // 用户名
	Code     string     `json:"code" example:"demo"`     // 用户编码（允许访问用户）默认=ID
	Summary  string     `json:"summary" example:"demo"`  // 简介
	ThirdUrl string     `json:"thirdUrl" example:"demo"` // 三方网站（需要审核）
	WaitUrl  string     `json:"waitUrl" example:"demo"`  // 待审核三方URL
	SourceId uint64     `json:"sourceId" example:"0"`    // 头像资源ID（仅小程序提供）
	OpenId   string     `json:"openId" example:"demo"`   // 小程序账号绑定登录（允许授权转移）
	Status   string     `json:"status" example:"0"`      // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"`                // 创建时间
	UpdateAt *time.Time `json:"updateAt"`                // 更新时间
}

// UserPageReq 极简用户信息 分页请求，根据实际业务替换分页条件字段
type UserPageReq struct {
	Id uint64 `json:"id" example:"1"` // 数据ID
	baseModel.PageReq
}

// UserPageRes 极简用户信息 分页响应，酌情从详情摘出部分字段
type UserPageRes struct {
	UserGetRes
}

// ToUserGetRes 极简用户信息 数据库转为详情响应
func ToUserGetRes(r *blogDB.User) *UserGetRes {
	return &UserGetRes{
		Id:       r.Id,
		ClientId: r.ClientId,
		Name:     r.Name,
		Code:     r.Code,
		Summary:  r.Summary,
		ThirdUrl: r.ThirdUrl,
		WaitUrl:  r.WaitUrl,
		SourceId: r.SourceId,
		OpenId:   r.OpenId,
		Status:   r.Status,
		CreateAt: r.CreateAt,
		UpdateAt: r.UpdateAt,
	}
}

// ToUserPageRes 极简用户信息 数据库转分页响应
func ToUserPageRes(list []*blogDB.User) []*UserPageRes {
	res := make([]*UserPageRes, len(list))
	for i, r := range list {
		res[i] = &UserPageRes{
			UserGetRes: *ToUserGetRes(r),
		}
	}
	return res
}
