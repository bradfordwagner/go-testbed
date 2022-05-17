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

	cron.New()
}
