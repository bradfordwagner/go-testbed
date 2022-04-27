package main

import (
	"errors"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()

	a := "a"
	l.Info("a: ", a)

	a, err := "b", errors.New("test error")
	l.WithError(err).Info("a: ", a)

	var f = func() {
		b, err := "1234", errors.New("func error")
		l.WithError(err).Info("b: ", b)
	}
	f()
	l.WithError(err).Info("a: ", a) // note the error set by func isn't in scope

}
