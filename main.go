package main

import "lafitd/runner"
import "github.com/sirupsen/logrus"
func main() {

	err := runner.Run()
	logrus.Info("Hello world")
	if err != nil {
		println(err.Error())
	}
}
