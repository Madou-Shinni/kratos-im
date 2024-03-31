package constants

// HandleResult 申请处理结果
type HandleResult int

const (
	HandleResultNone   HandleResult = iota + 1 // 未处理
	HandleResultRefuse                         // 拒绝
	HandleResultIgnore                         // 忽略
	HandleResultAgree                          // 同意
	HandleResultCancel                         // 撤销
)

// GroupRoleLevel 群等级 1. 创建者，2. 管理者，3. 普通
type GroupRoleLevel int

const (
	CreatorGroupRoleLevel GroupRoleLevel = iota + 1 // 为什么会 从1开始？
	ManagerGroupRoleLevel
	AtLargeGroupRoleLevel
)

// GroupJoinSource 进群申请的方式： 1. 邀请， 2. 申请
type GroupJoinSource int

const (
	InviteGroupJoinSource GroupJoinSource = iota + 1
	PutInGroupJoinSource
)
