package task

import (
	"context"
	v1 "task-schedule/api/task/v1"
)

type Controller struct {
}

func (c Controller) Submit(ctx context.Context, req *v1.TaskSubmitReq) (res *v1.TaskSubmitRes, err error) {
	//TODO implement me
	panic("implement me")
}

func (c Controller) Detail(ctx context.Context, req *v1.TaskDetailReq) (res *v1.TaskDetailRes, err error) {
	//TODO implement me
	panic("implement me")
}

func NewController() *Controller {
	return &Controller{}
}
