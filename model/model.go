package model

import (
	"fmt"
	"time"
)

type CoreOpt struct {
	Crontab   string        `clop:"short;long" usage:"crontab" default:"* * * * * *"`
	Count     int           `clop:"long" usage:"run count"`
	Output    string        `clop:"long" usage:"output"`
	Durations time.Duration `clop:"short;long" usage:"duration" default:"10s"`
}

func (c *CoreOpt) OutputFunc() {
	if len(c.Output) > 0 {
		fmt.Println(c.Output)
	}
}

func (c *CoreOpt) Sleep() {
	time.Sleep(c.Durations)
}
