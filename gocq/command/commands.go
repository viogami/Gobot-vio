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
}

var CommandList = []Command{
	NewCmdNull(),
	NewCmdHelp(),
	NewCmdSetu(),
	NewCmdSetuR18(),
	NewCmdChat(),
	NewCmdHuntSound(),
	NewCmdHuntSoundList(),
	NewCmdBanLottery(),
	NewCmdSetManager(),
	NewCmdGetRecall(),
}

var CommandMap = map[string]Command{
	"":       NewCmdNull(),
	"/help":  NewCmdHelp(),
	"/涩图":    NewCmdSetu(),
	"/涩图r18": NewCmdSetuR18(),
	"/chat":  NewCmdChat(),
	"/枪声":    NewCmdHuntSound(),
	"/枪声目录":  NewCmdHuntSoundList(),
	"/禁言抽奖":  NewCmdBanLottery(),
	"/给我管理":  NewCmdSetManager(),
	"/撤回了什么": NewCmdGetRecall(),
}
