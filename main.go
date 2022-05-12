package main

import "github.com/sirupsen/logrus"

func main() {
	l := logrus.New()

	m := map[string]string{
		"a": "a",
	}

	l.Info("map: ", m)

	delete(m, "b")

	l.Info("map: ", m)

}
