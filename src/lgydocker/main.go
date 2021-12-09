package main

import (
	log "github.com/sirupsen/logrus"//这里把logrus这个库自定为log
	"github.com/urfave/cli"//定义命令行参数的工具库
	"os"
)

const usage = "This is a simple implemetation of docker"

func main() {
	app := cli.NewApp()
	app.Name = "lgydocker"
	app.Usage = usage

	app.Commands = []cli.Command {
		initCommand,//这两个命令项是自定义的命令行命令
		runCommand,
	}

	app.Before = func(context *cli.Context) error {//这是cli工具库的初始化函数
		//使用json格式作为日志格式
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
