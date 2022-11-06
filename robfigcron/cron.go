package robfigcron

import (

	//"github.com/jakecoffman/cron"
	"github.com/guonaihong/crontest/model"
	cron "github.com/robfig/cron/v3"
)

type Robfigcron struct {
	model.CoreOpt
}

func (r *Robfigcron) SubMain() {

	c := cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))

	c.Start()

	var err error
	for i := 0; i < r.Count; i++ {
		_, err = c.AddFunc(r.Crontab, func() { r.Func() })
		if err != nil {
			panic(err.Error())
		}
	}

	r.Sleep()
}
