package gocq

// 用户权限
const (
	PermissionMember  = 0 // 用户权限
	PermissionAdmin = 1 // 管理员权限
	PermissionOwner = 2 // 群主权限
)

// 鉴权
func WhoYouAre(userID int64, groupID int64) int {
	// TODO
	return PermissionMember
}
