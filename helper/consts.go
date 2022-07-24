package helper

const (
	ServiceHttp = 1
	ServiceRPC  = 2
	ServiceTask = 3
)

const (
	Failure                        = "2000" //失败
	Pending                        = "2001" //待处理
	Processing                     = "2002" //处理中
	Success                        = "1000" //成功
	Degrade                        = "1001" //被降级过滤的请求
	AreaLimit                      = "1002" //地理区域限制
	AccessTokenExpires             = "1003" //Token过期
	UserDuplicate                  = "1004" //重复的用户
	UserLevelLow                   = "1005" //用户等级太低
	Blocked                        = "1006" //被锁定
	UsernameOrPasswordErr          = "1007" //用户名或密码错误
	PasswordTooLeak                = "1008" //密码太弱
	UserNotExist                   = "1009" //用户不存在
	FailedTooManyTimes             = "1010" //登录失败次数太多
	FileTooLarge                   = "1011" //上传文件太大
	LimitExceed                    = "1012" //超出限制
	ServerErr                      = "1013" //服务器错误
	UserIDCheckInvalidCard         = "1014" //请先完成实名认证
	UserIDCheckInvalidPhone        = "1015" //请先绑定手机
	UserIDCheckInvalid             = "1016" //账号尚未实名认证
	ServiceUpdate                  = "1017" //系统升级中
	MobileNoVerfiy                 = "1018" //未绑定手机
	CaptchaErr                     = "1019" //验证码错误
	UserDisabled                   = "1020" //账号被封停
	LackOfBalance                  = "1021" //余额不足
	MethodNoPermission             = "1022" //没有权限
	AmountInsufficient             = "1023" //场馆余额不足
	ParamNull                      = "1024" //参数为空
	ParamErr                       = "1025" //参数错误
	OldPasswordErr                 = "1026" //原密码错误
	DeviceTypeErr                  = "1027" //设备类型错误
	RedisErr                       = "1028" //Redis错误
	PasswordInconsistent           = "1029" //两次密码不一致
	DBErr                          = "1030" //数据库错误
	UsernameErr                    = "1031" //用户名错误
	PlatformMaintain               = "1032" //场馆维护中
	TransferOutErr                 = "1033" //场馆转出错误
	TransferInErr                  = "1034" //场馆转入错误
	BalanceErr                     = "1035" //查询余额错误
	OrderStateErr                  = "1036" //订单状态异常
	PlatformNotExist               = "1037" //场馆不存在
	FileTypeErr                    = "1038" //文件格式不合法
	AmountOutRange                 = "1039" //额度超过最大限制
	SessionErr                     = "1040" //session设置错误
	GetRPCErr                      = "1041" //rpc获取信息错误
	UpdateRPCErr                   = "1042" //rpc更新信息失败
	TransErr                       = "1043" //事务执行失败
	OrderProcess                   = "1044" //有处理中的转账订单
	PlatformNotTransferIn          = "1045" //场馆没有转入记录
	UserIDCheckInvalidEmail        = "1046" //请先绑定邮箱
	PhoneVerificationErr           = "1047" //手机验证码错误
	EmailVerificationErr           = "1048" //邮箱验证码错误
	CaptchaNull                    = "1049" //验证码不能为空
	PhoneExist                     = "1050" //手机号码已存在
	EmailExist                     = "1051" //邮箱已存在
	BankCardExistErr               = "1052" //银行卡已存在
	GenderAlreadyBind              = "1053" //性别已绑定
	BirthAlreadyBind               = "1054" //生日已绑定
	RealNameAlreadyBind            = "1055" //真实姓名已绑定
	NoDataUpdate                   = "1056" //无数据更新
	AmountErr                      = "1057" //金额错误
	RequestBusy                    = "1058" //请求频繁
	ESErr                          = "1059" //ES查询错误
	BankCardNotExist               = "1060" //银行卡不存在
	CateNotExist                   = "1061" //渠道不存在
	UsernameExist                  = "1062" //用户名已存在
	ChannelExist                   = "1063" //通道已存在
	CateExist                      = "1064" //渠道已存在
	ChannelNotExist                = "1065" //通道不存在
	OrderExist                     = "1066" //订单已存在
	DepositFailure                 = "1067" //支付失败
	WithdrawFailure                = "1068" //出款失败
	OrderNotExist                  = "1069" //订单不存在
	FormatErr                      = "1070" //格式解析错误
	MaxThreeBankCard               = "1071" //最多绑定3张银行卡
	PlatformRegErr                 = "1072" //场馆注册失败
	WaterFlowUnreached             = "1073" //流水未达标
	VerificationNumberNull         = "1074" //验证码次数已用完
	VerificationSendErr            = "1075" //验证码发送失败
	RequestFail                    = "1076" //请求失败
	PlatIDErr                      = "1077" //场馆ID错误
	AdjustModeErr                  = "1078" //调整模式错误
	TurnoverMutiErr                = "1079" //流水倍数错误
	ImagesURLErr                   = "1080" //图片地址错误
	UIDErr                         = "1081" //UID错误
	ApplyNameErr                   = "1082" //申请人错误
	ReviewNameErr                  = "1083" //审核人错误
	ReviewStateErr                 = "1084" //审核状态错误
	AdjustTyErr                    = "1085" //调整类型错误
	HandOutStateErr                = "1086" //发放状态错误
	SeamoErr                       = "1087" //seamo错误
	StateParamErr                  = "1088" //状态参数错误
	GroupIDErr                     = "1089" //用户分组错误
	IDErr                          = "1090" //ID错误
	PasswordFMTErr                 = "1091" //密码格式错误
	IPErr                          = "1092" //ip错误
	VersionErr                     = "1093" //version错误
	URLErr                         = "1094" //URL错误
	RemarkFMTErr                   = "1095" //备注格式错误
	BankcardIDErr                  = "1096" //银行账户错误
	DateTimeErr                    = "1097" //时间格式错误
	DeviceErr                      = "1098" //设备错误
	FileNotExist                   = "1099" //文件不存在
	WalletTypeErr                  = "1100" //钱包类型错误
	DividendTypeErr                = "1101" //红利类型错误
	PhoneFMTErr                    = "1102" //手机号格式错误
	AgentNameErr                   = "1103" //代理名错误
	EmailFMTErr                    = "1104" //邮件格式错误
	UserTagErr                     = "1105" //会员标签错误
	GenderErr                      = "1106" //性别错误
	ContentLengthErr               = "1107" //文本长度错误
	FileURLErr                     = "1108" //文件路径错误
	RedirectURLErr                 = "1109" //跳转URL格式错误
	PushDeviceErr                  = "1110" //推送设备类型错误
	MemberLevelErr                 = "1111" //会员等级错误
	ModuleErr                      = "1112" //应用模块错误
	CommonFlagErr                  = "1113" //是否常用标志错误
	ScenesErr                      = "1114" //应用场景错误
	PlatSeqErr                     = "1115" //场馆排序值错误
	PlatWalletErr                  = "1116" //场馆钱包值错误
	PLatNameErr                    = "1117" //场馆名错误
	GameTypeErr                    = "1118" //游戏类型错误
	MemberLevelAlterTyErr          = "1119" //会员等级调整类型错误
	CreateNameErr                  = "1120" //创建用户名错误
	RecordExistErr                 = "1121" //记录已存在
	RecordNotExistErr              = "1122" //记录不存在
	MaxDrawLimitParamErr           = "1123" //单次提款限额必须小于等于单日提款限额
	AgentTypeErr                   = "1124" //代理类型错误
	PresettleFlagErr               = "1125" //提前结算标志错误
	NoteNotNoneErr                 = "1126" //备注不能为空
	RecordIDErr                    = "1127" //记录ID错误
	RealNameFMTErr                 = "1128" //真实姓名错误
	OperateFailed                  = "1129" //操作失败
	GetDataFailed                  = "1130" //获取数据失败
	QueryTimeRangeErr              = "1131" //查询时间范围错误
	BlockchainProtocolErr          = "1132" //虚拟币协议错误
	VirtualWalletAddressExist      = "1133" //虚拟币地址重复
	VirtualWalletUpperLimit        = "1134" //虚拟币钱包数目超过上限
	TradeTypeErr                   = "1135" //交易类型错误
	FileTypeFlagErr                = "1136" //文件类型标志错误
	StaticFileUploadFailed         = "1137" //静态文件上传失败
	FileOpenFailed                 = "1138" //文件打开失败
	QueryTermsErr                  = "1139" //查询条件错误
	UnfreezeTyErr                  = "1140" //解除限制类型错误
	PhoneBindAlreadyErr            = "1141" //已经绑定手机，请勿重复绑定
	CashTypeErr                    = "1142" //帐变类型错误
	TransferTypeErr                = "1143" //场馆转账类型错误
	BetAmountRangeErr              = "1144" //投注额查询范围错误
	MemberTagInUse                 = "1145" //该标签会员正在使用中
	MerchantIDErr                  = "1146" //商户ID错误
	CateNameErr                    = "1147" //渠道名错误
	PromoApplyProcessing           = "1148" //活动申请处理中
	AtMost3MatchesOpen             = "1149" //最多允许开启三个赛事
	DeleteMatchNeedCloseFirst      = "1150" //删除赛事前请先关闭
	NoDepositErr                   = "1151" //未充值
	PromoTypeErr                   = "1152" //活动类型错误
	PromoGiftTypeErr               = "1153" //活动礼品类型错误
	ShipNameErr                    = "1154" //收货人名错误
	ShipAddressErr                 = "1155" //收货地址错误
	GiftNameErr                    = "1156" //礼品名错误
	GiftAttributes                 = "1157" //礼品属性错误
	GiftApplyTypeErr               = "1158" //礼品申请类型错误
	DailyAmountLimitErr            = "1159" //当日提款金额已上限，请于明日提款
	DailyTimesLimitErr             = "1160" //当日提款次数已达上限，请于明日提款
	CodeErr                        = "1161" //状态码错误
	OIDErr                         = "1162" //OID错误
	PasswordConsistent             = "1163" //新密码不能和原密码一致
	ChannelIDErr                   = "1164" //通道ID错误
	UsernamePhoneMismatch          = "1165" //手机号和用户名不匹配
	UsernameEmailMismatch          = "1166" //邮箱和用户名不匹配
	AdminNameErr                   = "1167" //操作用户名错误
	FinanceTypeErr                 = "1168" //财务类型错误
	TimeTypeErr                    = "1169" //筛选时间类型错误
	TunnelMinLimitErr              = "1170" //最小限额错误
	TunnelMaxLimitErr              = "1171" //最大限额错误
	OnlineCateDeleteErr            = "1172" //正在使用的渠道不能删除
	CateHaveChannelDeleteErr       = "1173" //渠道下级有通道就不能删除
	MerchantIDOrCateNameExist      = "1174" //商户ID或渠道名已存在
	CateIDAndChannelIDErr          = "1175" //渠道ID或通道ID错误
	BankNameOrCodeErr              = "1176" //银行编码或银行名错误
	InviteUsernameErr              = "1177" //邀请人用户名错误
	UpdateMustCloseFirst           = "1178" //请关闭后进行更新
	DeleteMustCloseFirst           = "1179" //请关闭后进行删除
	VerificationCodeNotExist       = "1180" //验证码不存在
	TunnelLimitParamErr            = "1181" //限额参数错误
	ParentChannelClosed            = "1182" //父通道已经关闭
	MemberLockAlready              = "1183" //会员已锁定
	OrderNumErr                    = "1184" //订单数目错误
	PromoIDErr                     = "1185" //活动id错误
	PromoClosed                    = "1186" //活动已关闭
	DuplicateApplyErr              = "1187" //已参加活动，不能重复申请
	GiftApplyTimeExpired           = "1188" //礼品兑换时间已过期
	EditMustCloseFirst             = "1189" //活动已开启,请关闭后在编辑
	PromoTitleErr                  = "1190" //活动自定义标题错误
	PromoValidDaysErr              = "1191" //活动兑换有效期天数错误
	PromoGrandPeriodErr            = "1192" //活动累计计数周期天数错误
	PlatNoPromoApply               = "1193" //该场馆未参加活动
	EmptyOrder5MinsBlock           = "1194" //由于您多次提交未存款订单，请于5分钟后再发起存款
	EmptyOrder5HoursBlock          = "1195" //由于您多次提交未存款订单，请于24小时后再发起存款
	NoChannelErr                   = "1196" //暂无存款通道
	MustLeaveAtLeastOneNotice      = "1197" //不能删除最后一条公告
	PromoApplyFailed               = "1198" //活动申请失败
	GroupExistAlready              = "1199" //团队已存在
	AgentPasswordErr               = "1200" //代理密码格式错误
	CanOnlyDealReviewing           = "1201" //只能审核或拒绝待审核的站内信
	CanOnlyOpenReviewing           = "1202" //只能启用待审核或停用的站内信
	OnlyOpenStopStatus             = "1203" //状态启用后才能停用
	ChannelBusyTryOthers           = "1204" //该通道拥挤，请尝试其他存款方式，或稍后再试
	MemberBankcardChannelUnsupport = "1205" //该通道暂不支持该会员提款银行卡
	NoPayChannel                   = "1206" //暂无可使用的代付通道
	DevelopNameErr                 = "1207" //发展人错误
	MaintainNameErr                = "1208" //维护人错误
	LinkExist                      = "1209" //维护人错误
	DynamicVerificationCodeErr     = "1210" //动态验证码错误
	CanOnlyJoinOneGroup            = "1211" //同一代理仅能组建一个团队
	SubGroupAlready                = "1212" //该代理已经是副线
	OperateNameErr                 = "1213" //操作人帐号错误
	TransferFromAgencyErr          = "1214" //转前代理帐号错误
	TransferToAgencyErr            = "1215" //转后代理帐号错误
	NoGroupErr                     = "1216" //您暂无团队
	OrderTakenAlready              = "1217" //该订单已被领取
	IsAgentUserAlready             = "1218" //该会员已经是%s代理所属会员
	AtMostAdd10MembersOnce         = "1219" //一次最多允许添加十个会员
	CanNotCancelGroupAgency        = "1220" //仅能取消不在团队中的代理资格，请将其移出团队后再进行此操作
	ManualPicking                  = "1221" //手动领取中，无法调整接单人员
	CanNotJoinOrCreateGroup        = "1222" //该代理本月无法加入或新建团队，请下月再试
	TypeCanNotJoinGroup            = "1223" //该类账号无法加入团队
	OfficialAgencyWithdrawNotAllow = "1224" //官方代理不允许提现
	CommissionRateLimitErr         = "1225" //佣金比例取值区间0-100
	DeviceBanErr                   = "1226" //该设备已经禁止注册，如有需要请联系客服。
	IpBanErr                       = "1227" //您的IP %s 已经禁止注册，如有需要请联系客服。
	UsernameFMTErr                 = "1228" //账号格式错误[4-9个字符，前2位必须为字母，数字可选，不支持符号]
	WithdrawBan                    = "1229" //暂无法提现，请联系客服
	WaterLimitUnReach              = "1230" //您的有效流水暂时不满足要求，请继续加油！
	SignAlreadyToday               = "1231" //今日已签到
	NoNeedResign                   = "1232" //没有漏签，无需补签
	BankcardBan                    = "1233" //该银行卡被禁止使用
	BankcardAbnormal               = "1234" //银行卡异常，请联系客服
	NoAwardCollect                 = "1235" //没有可领取的奖励
	WeeklyResignExpire             = "1236" //每周只能补签一次
	DisbandGroupErr                = "1237" //请先结算完本月以前所有佣金后再解散团队
	WithDrawProcessing             = "1238" //有正在提款中的银行卡无法删除，请稍候再试
	NotRedemptionTime              = "1239" //兑换时间未到,暂不可兑换奖品
	EmptyOrder30MinsBlock          = "1240" //由于您多次提交未存款订单，请于30分钟后再发起存款
	SingleWalletUnlockErr          = "1241" //您当前%s场馆已完成流水为%s,需达到%s才能解锁!
	ExchangeRateRrr                = "1242" //汇率错误
	OnlyOneBankcardActivePerBank   = "1243" //同一银行只能开启一张收款银行卡
	ChangeDepositLimitBeforeActive = "1244" //该银行卡今日收款限额已满，请修改今日收款限额后开启
	InvalidTransactionHash         = "1245" // 交易hash不合法
	ParentAgencyCanNotUse          = "1246" //上级已禁用 无法开启
	RebateOutOfRange               = "1247" //返水范围错误
	WithdrawPwdExist               = "1248" //提款密码重复设置
	SetWithdrawPwdFirst            = "1249" //请先设置提款密码
	WPwdCanNotSameWithPwd          = "1250" //提款密码不能和登录密码相同
	WithdrawPwdMismatch            = "1251" //提款密码错误
	NotDirectSubordinate           = "1252" //不是直属下级
	UsedCoPlanEditNotAllow         = "1253" //使用中的佣金方案不允许修改
	BallQuotaOutOfRange            = "1254" //单号码额度不足
	VncpIssueExpires               = "1255" //彩票期号过期
	MemberHaveSubAlready           = "1256" //当前会员已有下级
	IsAgentSubAlready              = "1257" //已经是当前代理下级
	TransferApplyExist             = "1258" //转代记录已存在
	RecordExpired                  = "1259" //站内信记录已过期
	ZaloExist                      = "1260" //zalo已存在
	AddressFMTErr                  = "1261" //收货地址格式错误
	ZaloBindAlreadyErr             = "1262" //已经绑定zalo，请勿重复绑定
	ZaloFMTErr                     = "1263" //zalo格式错误
	PayServerErr                   = "1264" //三方支付服务请求失败
	LoseAmountUnreached            = "1265" //负盈利未达标
	DepositAmountUnreached         = "1266" //存款未达标
	ApplyTurnErr                   = "1267" //活动申请顺序错误
	PreTaskUnFinishErr             = "1268" //您未完成流水活动，请完成再申请
	BankcardValidErr               = "1269" //银行卡真伪验证错误
	ThisTypeCanOnlyOpenOneErr      = "1270" //同类活动同时只能开启一个
	RegLimitExceed                 = "1271" //单设备注册超过最大值
	PromoExpired                   = "1272" //活动已过期
	FirstDailyWithdrawNeedVerify   = "1273" //每日第一笔提款需要短信验证
	MemberRebateModDisable         = "1274" //禁止编辑代理返水比例
	SubPermissionEqualErr          = "1275" //上下级权限相同
	MustApplyAfter1AM              = "1276" //请下在1点以后再申请
	DeleteBankcardInBlackListErr   = "1277" //不允许删除黑名单中的银行卡
	NotAllowAddLinkErr             = "1278" //不允许新增推广链接
	NotAllowDeleteLinkErr          = "1279" //不允许删除推广链接
	NotAllowModifySubRebateErr     = "1280" //不允许修改下级返水/返点比例
	PlatformLoginErr               = "1281" //场馆登陆失败
	PlatformTransferInErr          = "1282" //场馆转入失败
	PlatformTransferOutErr         = "1283" //场馆转出失败
	PlatformTransferCheckErr       = "1284" //场馆转账查单失败
	PlatformTransferInSuccess      = "1285" //场馆转入成功
	PlatformTransferInFailed       = "1286" //场馆转入失败
	PlatformTransferOutSuccess     = "1287" //场馆转出成功
	PlatformTransferOutFailed      = "1288" //场馆转出失败
)
