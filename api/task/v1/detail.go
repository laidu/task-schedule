package v1

import "github.com/gogf/gf/v2/frame/g"

type TaskDetailReq struct {
	g.Meta `path:"/detail" tags:"Task" method:"post" summary:"任务详情"`
	TaskId string `json:"task_id" binding:"required"`
}
type TaskDetailRes struct {
	g.Meta `mime:"text/json"`
	TaskId string `json:"task_id"`
}
