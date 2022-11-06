package antlabscronex

import (
	"fmt"
	"time"

	"github.com/antlabs/cronex"
	"github.com/guonaihong/crontest/model"
)

type AntlabsCronex struct {
	model.CoreOpt
}

func (r *AntlabsCronex) SubMain() {

	c := cronex.New()

	c.Start()

	var err error
	for i := 0; i < r.Count; i++ {
		_, err = c.AddFunc("* * * * *", func() { fmt.Println("Every Second") })
		fmt.Println(err)
	}

	// Inspect the cron job entries' next and previous run times.
	//inspect(c.Entries())
	//c.Stop() // Stop the s
	time.Sleep(time.Second * 1000)
	c.Stop()
}
