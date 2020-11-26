package cached

// Admin 存储当前登录用户信息
var Admin CurrentAdmin

// CurrentAdmin 存储当前登录用户信息
type CurrentAdmin struct {
	ID       uint
	UserName string
}

// Save 存储
func Save(token *Claims) {
	Admin.ID = token.ID
	Admin.UserName = token.UserName
}
