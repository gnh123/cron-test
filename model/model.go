package model

import "fmt"

type CoreOpt struct {
	Crontab string `clop:"short;long" usage:"crontab"`
	Count   int    `clop:"long" usage:"run count"`
	Output  string `clop:"long" usage:"output"`
}

func (c *CoreOpt) OutputFunc() {
	if len(c.Output) > 0 {
		fmt.Println(c.Output)
	}
}
