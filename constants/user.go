package constants

const (
	RedisKeySystemRootUid = "kratos-im:root"        // 内部root
	RedisKeyOnlineUser    = "kratos-im:online:user" // 在线用户key
	RedisKeyDiscoverSvr   = "kratos-im:server"      // 服务发现
)

type LoginType int32

const (
	LoginTypeAccount = iota + 1 // 账号(或邮箱)密码登录
	LoginTypeGithub             // github登录
)
