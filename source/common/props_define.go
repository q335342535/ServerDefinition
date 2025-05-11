package common

const (
	PropType_Count = 1
	PropType_Time  = 2
)

const (
	PropID_Score                 = 1
	PropID_LoongYearHead_Forever = 2
	PropID_LoongYearHead_7Days   = 3
)

//分数10000,PropStatus=0,分数是PropCount=10000，分数的TotalTerm是-1，表示永久，LeaveTerm是-1，表示永久
//3个7日记牌器，7日棋牌器是PropCount=3，7日棋牌器的TotalTerm是7*24*60*60，LeaveTerm是7*24*60*60，表示7天
//1个7日记牌器(14日过期),1个7日记牌器(3日过期)，1个7日记牌器(2025-05-09过期)
//以龙年头像为例，龙年头像的PropCount=1，龙年头像的TotalTerm是-1，表示永久，LeaveTerm是-1，表示永久

//道具未使用Redis，道具使用中Redis
//道具已使用表

type Prop struct {
	PropID        uint32 //道具ID
	PropStatus    uint8  //道具状态	无状态0、未使用、使用中、已使用、已过期
	PropType      uint8  //道具类型
	PropName      string //道具名称
	PropCount     uint64 //道具
	PropTotalTerm int64  //-1表示永久
	PropLeaveTerm int64  //-1表示永久
}
