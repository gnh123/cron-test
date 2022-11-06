package antlabscronex

import (
	"fmt"

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
		_, err = c.AddFunc(r.Crontab, func() { fmt.Println("Every Second") })
		fmt.Println(err)
	}

	r.Sleep()
	c.Stop()
}
