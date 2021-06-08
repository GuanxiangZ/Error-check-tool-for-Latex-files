package main

import (
	"github.com/sirupsen/logrus"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/command"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	command.Execute()
}

