package antlabscronex

import (
	"math/rand"
	"time"

	"github.com/antlabs/cronex"
	"github.com/antlabs/timer"
	"github.com/guonaihong/crontest/model"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type AntlabsCronex struct {
	model.CoreOpt
}

func (r *AntlabsCronex) SubMain() {

	c := cronex.New()

	c.Start()

	var err error
	allNodes := make([]cronex.TimerNoder, 0, r.Count)
	var tm timer.TimeNoder
	for i := 0; i < r.Count; i++ {
		tm, err = c.AddFunc(r.Crontab, func() { r.Func() })
		if err != nil {
			panic(err.Error())
		}
		allNodes = append(allNodes, tm)
	}

	go func() {
		var err error
		for {
			for i := 0; i < 100; i++ {

				index := rand.Int31n(int32(len(allNodes)))
				allNodes[index].Stop()
				allNodes[index], err = c.AddFunc(r.Crontab, func() { r.Func() })
				if err != nil {
					panic(err.Error())
				}
			}
			time.Sleep(time.Millisecond * 1000)
		}
	}()
	r.Sleep()
	c.Stop()
}
