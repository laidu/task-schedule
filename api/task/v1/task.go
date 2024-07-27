package v1

import "github.com/gogf/gf/v2/frame/g"

type TaskSubmitReq struct {
	g.Meta `path:"/submit" tags:"Task" method:"post" summary:"任务提交请求"`
}
type TaskSubmitRes struct {
	g.Meta `mime:"text/json"`
	TaskId string `json:"task_id"`
}

type TaskDetailReq struct {
	g.Meta `path:"/submit" tags:"Task" method:"post" summary:"任务提交请求"`
	TaskId string `json:"task_id" binding:"required"`
}
type TaskDetailRes struct {
	g.Meta `mime:"text/json"`
	TaskId string `json:"task_id"`
}
