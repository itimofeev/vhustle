package components

import (
	"github.com/itimofeev/vhustle/modules/gsheet"
	"github.com/robfig/cron"
)

func InitCronTasks() {
	c := cron.New()
	c.AddFunc("@every 1h", gsheet.UpdateContestsCron())
	c.Start()
}
