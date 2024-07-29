// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package task

import (
	"context"

	"task-scheduler/api/task/v1"
)

type ITaskV1 interface {
	TaskDetail(ctx context.Context, req *v1.TaskDetailReq) (res *v1.TaskDetailRes, err error)
	TaskSubmit(ctx context.Context, req *v1.TaskSubmitReq) (res *v1.TaskSubmitRes, err error)
}
