package main

import (
	_ "task-schedule/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"task-schedule/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
