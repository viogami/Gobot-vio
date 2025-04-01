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
	// newCmdHuntSoundList(),
	newCmdBanLottery(),
	newCmdSetManager(),
	newCmdGetRecall(),
}

var CommandMap = map[string]Command{
	"":       newCmdNull(),
	"/help":  newCmdHelp(),
	"/涩图":    newCmdSetu(),
	"/涩图r18": newCmdSetuR18(),
	"/chat":  newCmdChat(),
	"/枪声":    newCmdHuntSound(),
	// "/枪声目录":  newCmdHuntSoundList(),
	"/禁言抽奖":  newCmdBanLottery(),
	"/给我管理":  newCmdSetManager(),
	"/撤回了什么": newCmdGetRecall(),
}
