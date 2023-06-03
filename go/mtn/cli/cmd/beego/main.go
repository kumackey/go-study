package main

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
	"time"
)

func main() {
	log := logs.NewLogger(10000)
	log.SetLogger(logs.AdapterMultiFile, `{
	"filename": "./foo.log",
	"daily": true,
	"maxlines": 10000,
	"rotate": true,
}`)
	for i := 0; i < 100; i++ {
		log.Info("Hello World!")
	}
}

func task1() {
	tk1 := task.NewTask("tk1", "0/3 * * * *", func(ctx context.Context) error {
		logs.Info("tk1")
		return nil
	})

	err := tk1.Run(context.Background())
	if err != nil {
		logs.Error(err)
	}

	task.AddTask("tk1", tk1)
	task.StartTask()
	defer task.StopTask()

	time.Sleep(12 * time.Second)
}
