package task

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"task-scheduler/api/task/v1"
)

func (c *ControllerV1) TaskSubmit(ctx context.Context, req *v1.TaskSubmitReq) (res *v1.TaskSubmitRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
