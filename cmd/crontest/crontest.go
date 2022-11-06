package main

import (
	"github.com/guonaihong/clop"
	"github.com/guonaihong/crontest/antlabscronex"
	"github.com/guonaihong/crontest/robfigcron"
)

type crontest struct {
	// 子命令, 会调用该结构的SubMain函数
	antlabscronex.AntlabsCronex `clop:"subcommand" usage:"github.com/antlabs/cronex"`
	// 子命令, 会调用该结构的SubMain函数
	robfigcron.Robfigcron `clop:"subcommand" usage:"github.com/robifg/cron"`
}

func main() {
	var c crontest
	clop.Bind(&c)
}
