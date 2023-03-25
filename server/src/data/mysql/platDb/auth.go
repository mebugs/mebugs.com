package platDb

// AuthUser 授权用户结构体
type AuthUser struct {
	ID         uint64   // 账号ID
	Account    string   // 账号
	PwdExpTime string   // 密码超期时间（修改后的90天）
	TenantId   uint64   // 租户ID
	Status     uint8    // 账号状态 1正常 2锁定 3封存
	Roles      []string // 角色集（该账号具备的角色）
	Permission []string // 权限集（该账号具备的权限）
	Routers    []string // 路由集（该账号可以访问的后端路由）
}
