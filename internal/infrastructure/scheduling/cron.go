package scheduling

import (
	"github.com/robfig/cron/v3"
)

type goCron struct {
	*cron.Cron
}

func NewCronInstance() Scheduling {
	return &goCron{cron.New()}
}
