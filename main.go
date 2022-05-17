package main

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	l := logrus.New()

	loc := time.UTC
	l.Info("locale: ", loc.String())

	c := cron.New(cron.WithLocation(loc))
	_, err := c.AddJob("* * * * *", job{})
	if err != nil {
		l.WithError(err).Error("failed to add cron")
	}
	c.Run()
}

type job struct {
}

func (j job) Run() {
	logrus.Info("doing the things!")
}

var _ cron.Job = (*job)(nil)
