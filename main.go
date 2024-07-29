package main

import (
	_ "task-scheduler/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"task-scheduler/internal/cmd"
)

func main() {

	initCtx := gctx.GetInitCtx()

	// 启动任务调度服务
	go cmd.MachineryServer.Run(initCtx)

	//  启动http 服务
	cmd.Main.Run(initCtx)

}
