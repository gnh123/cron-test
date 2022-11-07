package robfigcron

import (

	//"github.com/jakecoffman/cron"
	"math/rand"
	"time"

	"github.com/guonaihong/crontest/model"
	cron "github.com/robfig/cron/v3"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Robfigcron struct {
	model.CoreOpt
}

func (r *Robfigcron) SubMain() {

	c := cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))

	c.Start()

	var err error
	allNodes := make([]cron.EntryID, 0, r.Count)
	var tm cron.EntryID
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
				id := allNodes[index]
				c.Remove(id)
				allNodes[index], err = c.AddFunc(r.Crontab, func() { r.Func() })
				if err != nil {
					panic(err.Error())
				}
			}
			time.Sleep(time.Millisecond * 1000)
		}
	}()
	r.Sleep()
}
