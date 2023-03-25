package platModel

// AuthLogin 账密登陆结构体
type AuthLogin struct {
	Account     string `json:"account" binding:"required,max=16"`        // 账号
	Password    string `json:"password" binding:"required,min=6,max=16"` // 密码
	TenantAlias string `json:"tenantAlias" binding:"required"`           // 租户别名
}
