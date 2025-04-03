package gocq

type GocqData struct {
	RecalledMsgStack []map[int64]recallMsgInfo
}

type recallMsgInfo struct {
	MsgId  int64
	OpId   int64
	UserId int64
}

// GocqDataInstance 是 GocqData 的单例,会在第一次调用NewGocqServer时初始化它
var GocqDataInstance *GocqData

func NewGocqData() *GocqData {
	gocqData := new(GocqData)
	gocqData.RecalledMsgStack = make([]map[int64]recallMsgInfo,0)

	return gocqData
}

func (g *GocqData) AddRecalledMsg(groupId int64, msgId int64, opid int64, userId int64) {
	gocqData := make(map[int64]recallMsgInfo)
	gocqData[groupId] = recallMsgInfo{
		MsgId:  msgId,
		OpId:   opid,
		UserId: userId,
	}
	g.RecalledMsgStack = append(g.RecalledMsgStack, gocqData)
}
func (g *GocqData) GetRecalledMsg(groupId int64) (int64, int64, int64) {
	if len(g.RecalledMsgStack) == 0 {
		return 0, 0, 0
	}
	for _, v := range g.RecalledMsgStack {
		if info, ok := v[groupId]; ok {
			return info.MsgId, info.OpId, info.UserId
		}
	}
	return 0, 0, 0
}
