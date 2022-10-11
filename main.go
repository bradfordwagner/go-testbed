package main

import (
	"errors"
	"github.com/avast/retry-go/v4"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	l := logrus.New()

	err := retry.Do(func() error {
		l.Info("hi")
		return errors.New("i am a bad error")
	}, retry.DelayType(retry.BackOffDelay), retry.MaxDelay(time.Second*5), retry.Attempts(5), retry.LastErrorOnly(true))

	if err != nil {
		l.WithError(err).Error("received error at end-step")
	}

}
