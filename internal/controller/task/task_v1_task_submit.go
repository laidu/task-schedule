package task

import (
	"context"
	"fmt"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"task-scheduler/internal/service"
	"time"

	"task-scheduler/api/task/v1"
)

func (c *ControllerV1) TaskSubmit(ctx context.Context, req *v1.TaskSubmitReq) (res *v1.TaskSubmitRes, err error) {

	machineryServer := service.NewMachineryServer()

	var (
		addTask0 tasks.Signature
	)

	taskArgs := make([]tasks.Arg, 0, len(req.Args))
	for _, arg := range req.Args {
		taskArgs = append(taskArgs, tasks.Arg{
			Name:  arg.Name,
			Type:  arg.Type,
			Value: arg.Value,
		})
	}

	addTask0 = tasks.Signature{
		Name: req.Name,
		Args: taskArgs,
	}

	/*
	 * Lets start a span representing this run of the `send` command and
	 * set a batch id as baggage so it can travel all the way into
	 * the worker functions.
	 */
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "send")
	defer span.Finish()

	batchID := uuid.New().String()
	span.SetBaggageItem("batch.id", batchID)
	span.LogFields(opentracinglog.String("batch.id", batchID))

	log.INFO.Println("Starting batch:", batchID)
	/*
	 * First, let's try sending a single task
	 */

	asyncResult, err := machineryServer.SendTask(&addTask0)
	if err != nil {
		return nil, fmt.Errorf("Could not send task: %s", err.Error())
	}

	results, err := asyncResult.Get(time.Millisecond * 5)
	if err != nil {
		return nil, err
	}

	return &v1.TaskSubmitRes{
		Meta:   g.Meta{},
		TaskId: addTask0.UUID,
		Result: results,
	}, nil
}
