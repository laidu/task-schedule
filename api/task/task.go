package task

import (
	"context"
	v1 "task-schedule/api/task/v1"
)

type ITaskV1 interface {

	// Submit 任务提交
	Submit(ctx context.Context, req *v1.TaskSubmitReq) (res *v1.TaskSubmitRes, err error)
	// Detail 任务详情
	Detail(ctx context.Context, req *v1.TaskDetailReq) (res *v1.TaskDetailRes, err error)
}
