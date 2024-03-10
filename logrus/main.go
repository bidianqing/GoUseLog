package main

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-extras/elogrus.v7"
)

func main() {
	logger := logrus.New()
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://127.0.0.1:9200"},
	})
	if err != nil {
		panic(err)
	}
	hook, err := elogrus.NewAsyncElasticHook(client, "localhost", logrus.InfoLevel, "golog")
	if err != nil {
		panic(err)
	}

	logger.Hooks.Add(hook)

	logger.SetLevel(logrus.InfoLevel)
	fmt.Println("日志级别为", logrus.GetLevel())
	fmt.Println("是否开启", logrus.ErrorLevel, logger.IsLevelEnabled(logrus.ErrorLevel))

	logger.Debug("Debug日志信息")
	logger.Info("Info日志信息")
	logger.Warn("Warn日志信息")
	logger.Error("Error日志信息")
}
