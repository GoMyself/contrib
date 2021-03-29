package helper


var ErrorCode = map[string]int{
	"Success"                 : 2000,
	"System"                  : 2001,   //系统错误,redis或者mysql异常抛出
	"Token"                   : 2002,   //Token 过期
	"UserDisabled"            : 2003,   //账号被封停
	"UserDuplicate"           : 2004,   //重复的用户
	"AccessDenied"            : 2005,   //访问权限不足
	"FailedTooManyTimes"      : 2006,   //登录失败次数太多
	"UsernameOrPasswordErr"   : 2007,   //用户名或密码错误
	"Deadline"                : 2008,   //服务调用超时
	"UserIDCheckInvalidPhone" : 2009,   //请先绑定手机
	"UserIDCheckInvalidCard"  : 2010,   //请先完成银行卡绑定
	"LackOfScores"            : 2011,   //积分不足
	"LackOfBalance"           : 2012,   //余额不足

}
