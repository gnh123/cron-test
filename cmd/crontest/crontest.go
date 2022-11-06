package main

import (
	"github.com/guonaihong/clop"
	"github.com/guonaihong/crontest/antlabscronex"
	"github.com/guonaihong/crontest/robfigcron"
)

type crontest struct {
	antlabscronex.AntlabsCronex `clop:"subcommand" usage:"github.com/antlabs/cronex"`
	robfigcron.Robfigcron       `clop:"subcommand" usage:"github.com/robifg/cron"`
}

func main() {
	var c crontest
	clop.Bind(&c)
}
