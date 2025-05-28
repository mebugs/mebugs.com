package constant

/**
 *
 * 响应码常量
 * 具体文言维护在数据库，但编码需要在此处维护，提高代码可读性
 *
 * @author 米虫丨www.mebugs.com
 * @since 2023-07-21
 */

const (
	CodeFmt = "%s%s%02d" // 成功码Format

	Error           = "E000" // 系统异常（默认）
	ValidErr        = "E001" // 参数非法（默认）（免翻译）
	LoginErr        = "E002" // 尚未登陆（默认）
	AuthErr         = "E003" // 无权访问（默认）
	PathErr         = "E004" // 路径不存在（默认）
	Success         = "S000" // 处理成功（默认）
	AuthLoginSS     = "S001" // 登陆成功
	AuthResetSS     = "S002" // 密码重置成功
	Fail            = "F000" // 处理失败（默认）
	AuthLoginNG     = "F001" // 登陆失败，请联系管理员
	AuthLoginRuleNG = "F002" // 异常登陆，请联系管理员
	AuthResetNG     = "F003" // 密码重置失败，请联系管理员

	DictAddSS  = "S100" // 字典创建成功
	DictEditSS = "S101" // 字典编辑成功
	DictSortSS = "S102" // 字典排序成功
	DictDelSS  = "S103" // 字典封存成功
	DictGetNG  = "F100" // 字典查询失败
	DictUniNG  = "F101" // 字典分组下字典值唯一
	DictMarkNG = "F102" // 内置字典禁止刪除
	DictSortNG = "F103" // 字典排序失败

	ResponseAddSS    = "S200" // 响应码创建成功
	ResponseAddSSWNC = "S201" // 响应码创建成功,实际响应码为{{code}}
	ResponseEditSS   = "S202" // 响应码编辑成功
	ResponseDelSS    = "S203" // 响应码封存成功
	ResponseGetNG    = "F200" // 响应码查询失败
	ResponseUniNG    = "F201" // 响应码全局唯一
	ResponseMarkNG   = "F202" // 内置响应码禁止删除

	RouterAddSS     = "S300" // 路由创建成功
	RouterEditSS    = "S301" // 路由编辑成功
	RouterDelSS     = "S302" // 路由删除成功
	RouterGetNG     = "F300" // 路由查询失败
	RouterUniUrlNG  = "F301" // 路由地址全局唯一
	RouterUniNameNG = "F302" // 路由名称全局唯一
	RouterMarkNG    = "F303" // 内置路由禁止删除
	RouterDelNG     = "F304" // 路由删除

	PermissionAddSS      = "S400" // 权限创建成功
	PermissionEditSS     = "S401" // 权限编辑成功
	PermissionDelSS      = "S402" // 权限删除成功
	PermissionSortSS     = "S403" // 权限排序成功
	PermissionGetNG      = "F400" // 权限查询失败
	PermissionUniAliasNG = "F401" // 权限别名全局唯一
	PermissionUniNameNG  = "F402" // 权限名称全局唯一
	PermissionMarkNG     = "F403" // 内置权限禁止删除
	PermissionRouterNG   = "F404" // 权限配置路由同步失败
	PermissionDelNG      = "F405" // 权限配置删除失败

	RoleAddSS        = "S500" // 角色创建成功
	RoleEditSS       = "S501" // 角色编辑成功
	RoleDelSS        = "S502" // 角色删除失败
	RoleGetNG        = "F500" // 角色查询失败
	RoleMarkNG       = "F501" // 内置角色禁止编辑
	RoleNameUniNG    = "F502" // 角色名全局唯一
	RolePermissionNG = "F503" // 角色权限配置失败
	RoleDelNG        = "F504" // 角色删除失败

	DeptAddSS        = "S600" // 集团部门创建成功
	DeptEditSS       = "S601" // 集团部门编辑成功
	DeptDelSS        = "S602" // 集团部门删除成功
	DeptSortSS       = "S603" // 集团部门排序成功
	DeptToSS         = "S604" // 集团部门迁移成功
	DeptGetNG        = "F600" // 集团部门查询失败
	DeptUniNameNG    = "F601" // 集团部门名称全局唯一
	DeptMarkNG       = "F602" // 内置集团部门禁止删除
	DeptDelNG        = "F603" // 集团部门删除失败
	DeptDelChildNG   = "F604" // 集团部门存在子部门禁止删除
	DeptDelAccountNG = "F605" // 集团部门存在成员禁止删除
	DeptToNG         = "F606" // 集团部门迁移失败

	AccountAddSS    = "S700" // 登陆账号创建成功
	AccountEditSS   = "S701" // 登陆账号编辑成功
	AccountDelSS    = "S702" // 登陆账号删除成功
	AccountResetSS  = "S703" // 登陆账号重置成功
	AccountGetNG    = "F700" // 登陆账号查询失败
	AccountUniXxxNG = "F701" // 登陆账号全局唯一
	AccountRoleNG   = "F702" // 账号角色同步失败
	AccountMarkNG   = "F703" // 特殊账号禁止编辑
	AccountDelNG    = "F704" // 登陆账号删除失败
	AccountResetNG  = "F705" // 登陆账号重置失败

	SysConfigEditSS = "S800" // 系统配置编辑成功
	SysConfigGetNG  = "F800" // 系统配置查询失败

	CategoryAddSS     = "S900" // 文章分类创建成功
	CategoryEditSS    = "S901" // 文章分类编辑成功
	CategoryDelSS     = "S902" // 文章分类删除成功
	CategoryMergeSS   = "S903" // 文章分类迁移成功
	CategoryGetNG     = "F900" // 文章分类查询失败
	CategoryMergeNG   = "F901" // 文章分类迁移失败
	CategoryAddNoScNG = "F902" // 文章分类未提交图片
	CategoryUniXxxNG  = "F903" // 文章分类地址全局唯一
	CategoryDelNoNG   = "F904" // 文章分类下存在数据禁止删除
	CategoryDelNG     = "F905" // 文章分类删除失败

	SourceAddSS    = "S1000" // 资源配置表创建成功
	SourceEditSS   = "S1001" // 资源配置表编辑成功
	SourceDelSS    = "S1002" // 资源清理暂不开放
	SourceGetNG    = "F1000" // 资源配置表查询失败
	SourceFileUpNg = "F1001" // 资源文件上传失败

	TagAddSS     = "S1100" // 标签创建成功
	TagEditSS    = "S1101" // 标签编辑成功
	TagDelSS     = "S1102" // 标签删除成功
	TagGetNG     = "F1100" // 标签查询失败
	TagAddNoScNG = "F1101" // 文章标签未提交图片
	TagUniUrlNG  = "F1102" // 标签地址全局唯一
	TagDelPostNG = "F1103" // 文章标签下数据删除失败
	TagDelNG     = "F1104" // 标签删除成功

	TopicAddSS     = "S1200" // 文章专题创建成功
	TopicEditSS    = "S1201" // 文章专题编辑成功
	TopicDelSS     = "S1202" // 文章专题删除成功
	TopicGetNG     = "F1200" // 文章专题查询失败
	TopicAddNoScNG = "F1201" // 文章标签未提交图片
	TopicUniInfoNG = "F1202" // 文章专题内文章应当唯一
	TopicPostNG    = "F1203" // 文章专题关联文章同步失败
	TopicDelNG     = "F1204" // 文章专题删除

	PostAddSS        = "S1300" // 文章或页面创建成功
	PostEditSS       = "S1301" // 文章或页面编辑成功
	PostGetNG        = "F1300" // 文章或页面查询失败
	PostAddNoScNG    = "F1301" // 文章或页面必须有主图
	PostTagNG        = "F1302" // 文章关联标签更新失败
	PostSourceNG     = "F1303" // 文章关联资源更新失败
	PostMoreNG       = "F1304" // 文章关联详情更新失败
	PostUniXxxNG     = "F1305" // 文章或页面地址全局唯一
	PostMoreUniXxxNG = "F1306" // 文章关联地址全局唯一

	BannerAddSS     = "S1400" // Banner创建成功
	BannerEditSS    = "S1401" // Banner编辑成功
	BannerDelSS     = "S1402" // Banner删除成功
	BannerGetNG     = "F1400" // Banner查询失败
	BannerAddNoScNG = "F1401" // Banner未提交图片
	BannerUniXxxNG  = "F1402" // Banner地址全局唯一
	BannerDelNG     = "F1403" // Banner删除失败

	PostCommentEditSS = "S1500" // 评论更新成功
	PostCommentGetNG  = "F1500" // 评论查询失败

	LinksAddSS     = "S1600" // 友情链接创建成功
	LinksEditSS    = "S1601" // 友情链接编辑成功
	LinksDelSS     = "S1602" // 友情链接删除成功
	LinksGetNG     = "F1600" // 友情链接查询失败
	LinksAddNoScNG = "F1601" // 友情链接未提交图片
	LinksUniXxxNG  = "F1602" // 友情链接地址全局唯一
	LinksDelNG     = "F1603" // 友情链接删除失败

)
