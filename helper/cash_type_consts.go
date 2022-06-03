package helper

// 帐变类型
const (
	TransactionIn                    = 151 //场馆转入
	TransactionOut                   = 152 //场馆转出
	TransactionInFail                = 153 //场馆转入失败补回
	TransactionOutFail               = 154 //场馆转出失败扣除
	TransactionDeposit               = 155 //存款
	TransactionWithDraw              = 156 //提现
	TransactionUpPoint               = 157 //后台上分
	TransactionDownPoint             = 158 //后台下分
	TransactionDownPointBack         = 159 //后台下分回退
	TransactionDividend              = 160 //中心钱包红利派发
	TransactionRebate                = 161 //会员返水
	TransactionFinanceDownPoint      = 162 //财务下分
	TransactionWithDrawFail          = 163 //提现失败
	TransactionValetDeposit          = 164 //代客充值
	TransactionValetWithdraw         = 165 //代客提款
	TransactionAgencyDeposit         = 166 //代理充值
	TransactionAgencyWithdraw        = 167 //代理提款
	TransactionPlatUpPoint           = 168 //后台场馆上分
	TransactionPlatDividend          = 169 //场馆红利派发
	TransactionSubRebate             = 170 //下级返水
	TransactionFirstDepositDividend  = 171 //首存活动红利
	TransactionInviteDividend        = 172 //邀请好友红利
	TransactionBet                   = 173 //投注
	TransactionBetCancel             = 174 //投注取消
	TransactionPayout                = 175 //派彩
	TransactionResettlePlus          = 176 //重新结算加币
	TransactionResettleDeduction     = 177 //重新结算减币
	TransactionCancelPayout          = 178 //取消派彩
	TransactionPromoPayout           = 179 //场馆活动派彩
	TransactionEBetTCPrize           = 600 //EBet宝箱奖金
	TransactionEBetLimitRp           = 601 //EBet限量红包
	TransactionEBetLuckyRp           = 602 //EBet幸运红包
	TransactionEBetMasterPayout      = 603 //EBet大赛派彩
	TransactionEBetMasterRegFee      = 604 //EBet大赛报名费
	TransactionEBetBetPrize          = 605 //EBet投注奖励
	TransactionEBetReward            = 606 //EBet打赏
	TransactionEBetMasterPrizeDeduct = 607 //EBet大赛奖金取回
	TransactionWMReward              = 608 //WM打赏
	TransactionSBODividend           = 609 //SBO红利
	TransactionSBOReward             = 610 //SBO打赏
	TransactionSBOBuyLiveCoin        = 611 //SBO 购买LiveCoin
	TransactionSignDividend          = 612 //天天签到活动红利
	TransactionCQ9Dividend           = 613 //CQ9游戏红利
	TransactionCQ9PromoPayout        = 614 //CQ9活动派彩
	TransactionPlayStarPrize         = 615 //Playstar积宝奖金
	TransactionSpadeGamingRp         = 616 //SpadeGaming红包
	TransactionAEReward              = 617 //AE打赏
	TransactionAECancelReward        = 618 //AE取消打赏
	TransactionOfflineDeposit        = 619 //线下转卡存款
	TransactionUSDTOfflineDeposit    = 620 //USDT线下存款
	TransactionEVOPrize              = 621 //游戏奖金(EVO)
	TransactionEVOPromote            = 622 //推广(EVO)
	TransactionEVOJackpot            = 623 //头奖(EVO)
	TransactionCommissionDraw        = 624 //佣金提取
	TransactionSettledBetCancel      = 625 //投注取消(已结算注单)
	TransactionCancelledBetRollback  = 626 //已取消注单回滚
	TransactionSBOReturnStake        = 627 //SBO ReturnStake
	TransactionBetNSettleWin         = 628 //电子投付赢
	TransactionBetNSettleLose        = 629 //电子投付输
	TransactionAdjustPlus            = 630 //场馆调整加
	TransactionAdjustDiv             = 631 //场馆调整减
	TransactionCQ9TakeAll            = 632 //CQ9捕鱼转入
	TransactionCQ9Refund             = 633 //CQ9游戏转出
	TransactionCQ9RollIn             = 634 //CQ9捕鱼转出
	TransactionCQ9RollOut            = 635 //CQ9捕鱼转入
	TransactionBetNSettleWinCancel   = 636 //投付赢取消
	TransactionBetNSettleLoseCancel  = 637 //投付输取消
	TransactionSetBalanceZero        = 638 //中心钱包余额冲正
	TransactionVIPUpgrade            = 639 //vip晋级礼金
	TransactionVIPMonthly            = 640 //vip月红包
	TransactionVIpBirthday           = 641 //vip生日礼金
	TransactionRebateCasino          = 642 //真人返水
	TransactionRebateLottery         = 643 //彩票返水
	TransactionRebateSport           = 644 //体育返水
	TransactionRebateDesk            = 645 //棋牌返水
	TransactionRebateESport          = 646 //电竞返水
	TransactionRebateCockFighting    = 647 //斗鸡返水
	TransactionRebateFishing         = 648 //捕鱼返水
	TransactionRebateLott            = 649 //电游返水
	TransactionRebateCGLottery       = 650 //彩票返点
)
