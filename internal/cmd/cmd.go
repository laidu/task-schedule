package cmd

import (
	"context"
	exampletasks "github.com/RichardKnop/machinery/v2/example/tasks"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
	"task-scheduler/internal/controller"
	"task-scheduler/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			controller.RegisterAPI(s)
			s.Run()
			return nil
		},
	}

	MachineryServer = gcmd.Command{
		Name:  "machinery",
		Usage: "machinery",
		Brief: "start machinery server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			machineryServer := service.NewMachineryServer()

			consumerTag := "machinery_worker"
			// The second argument is a consumer tag
			// Ideally, each worker should have a unique tag (worker1, worker2 etc)
			worker := machineryServer.NewWorker(consumerTag, 0)

			// Here we inject some custom code for error handling,
			// start and end of task hooks, useful for metrics for example.
			errorHandler := func(err error) {
				log.ERROR.Println("I am an error handler:", err)
			}

			preTaskHandler := func(signature *tasks.Signature) {
				log.INFO.Println("I am a start of task handler for:", signature.Name)
			}

			postTaskHandler := func(signature *tasks.Signature) {
				log.INFO.Println("I am an end of task handler for:", signature.Name)
			}

			worker.SetPostTaskHandler(postTaskHandler)
			worker.SetErrorHandler(errorHandler)
			worker.SetPreTaskHandler(preTaskHandler)

			// Register tasks
			tasksMap := map[string]interface{}{
				"add":               exampletasks.Add,
				"multiply":          exampletasks.Multiply,
				"sum_ints":          exampletasks.SumInts,
				"sum_floats":        exampletasks.SumFloats,
				"concat":            exampletasks.Concat,
				"split":             exampletasks.Split,
				"panic_task":        exampletasks.PanicTask,
				"long_running_task": exampletasks.LongRunningTask,
			}

			err = machineryServer.RegisterTasks(tasksMap)
			if err != nil {
				return err
			}

			return worker.Launch()
		},
	}
)
