package task

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"task-scheduler/api/task/v1"
)

func (c *ControllerV1) TaskDetail(ctx context.Context, req *v1.TaskDetailReq) (res *v1.TaskDetailRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
