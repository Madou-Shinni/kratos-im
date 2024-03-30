package constants

type HandleResult int

const (
	HandleResultNone   HandleResult = iota + 1 // 未处理
	HandleResultRefuse                         // 拒绝
	HandleResultIgnore                         // 忽略
	HandleResultAgree                          // 同意
	HandleResultCancel                         // 撤销
)
