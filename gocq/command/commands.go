package command

type Command interface {
	Execute(params CommandParams) // 执行指令
	GetInfo(index int) string     // 获取指令信息
}

type CommandParams struct {
	MessageId   int32
	MessageType string
	Message     string
	GroupId     int64
	UserId      int64

	SetuParams
}

type SetuParams struct {
	Tags []string // 标签
}

var CommandList = []Command{
	newCmdNull(),
	newCmdHelp(),
	newCmdSetu(),
	newCmdSetuR18(),
	newCmdChat(),
	newCmdHuntSound(),
	newCmdBanLottery(),
	newCmdGetRecall(),
}
var CommandMap = map[string]Command{
	"help":    newCmdHelp(),
	"来份涩图":    newCmdSetu(),
	"来份r18涩图": newCmdSetuR18(),
	"/chat":   newCmdChat(),
	"猎杀对决枪声":  newCmdHuntSound(),
	"禁言抽奖":    newCmdBanLottery(),
	"撤回了什么":   newCmdGetRecall(),
}

// var OriginalCommandMap = map[string]Command{
// 	"":       newCmdNull(),
// 	"/help":  newCmdHelp(),
// 	"/涩图":    newCmdSetu(),
// 	"/涩图r18": newCmdSetuR18(),
// 	"/chat":  newCmdChat(),
// 	"/枪声":    newCmdHuntSound(),
// 	"/禁言抽奖":  newCmdBanLottery(),
// 	"/撤回了什么": newCmdGetRecall(),
// }
