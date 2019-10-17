package main

import "flag"

type Command struct {
	UsageLine string
	Run       func(args []string) int
	Flag      flag.FlagSet
}

func (c *Command) Name() string {
	return c.UsageLine
}

var (
	trainTableCommand = &Command{
		UsageLine: "schedule",
	}
	trainInfoTableCommand = &Command{
		UsageLine: "info",
	}
	Commands = []*Command{
		trainTableCommand,
		trainInfoTableCommand,
	}

	trainTableFromStation string
	trainTableToStation   string
	trainTableDate        string

	trainInfoCode string
	trainInfoDate string
)

func executeTrainTableFunc(args []string) int {
	// 查询车次列表
	trainTable := NewTrainTable(args[0], args[1], args[2], "")
	trainTable.runParser()
	return 1
}

func executeTrainInfoFunc(args []string) int {
	// 查询车次信息
	trainInfo := NewTrainInfoTable(args[0], args[1])
	trainInfo.runParser()
	return 1
}

func commandLineInit() {
	trainTableCommand.Run = executeTrainTableFunc
	trainTableCommand.Flag.StringVar(&trainTableFromStation, "from", "", "需要搜索的起始站点")
	trainTableCommand.Flag.StringVar(&trainTableToStation, "to", "", "需要搜索的到达站点")
	trainTableCommand.Flag.StringVar(&trainTableDate, "date", "", "需要搜索的日期（格式: YYYY-MM-DD 例如: 2019-10-17）")

	trainInfoTableCommand.Run = executeTrainInfoFunc
	trainInfoTableCommand.Flag.StringVar(&trainInfoCode, "code", "", "需要搜索的车次号（例如：G1）")
	trainInfoTableCommand.Flag.StringVar(&trainInfoDate, "date", "", "需要搜索的日期（格式: YYYY-MM-DD 例如: 2019-10-17）")
}
