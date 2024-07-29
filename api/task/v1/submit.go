package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Arg struct {
	Name  string      `bson:"name"`
	Type  string      `bson:"type"`
	Value interface{} `bson:"value"`
}

type TaskSubmitReq struct {
	g.Meta `path:"/submit" tags:"Task" method:"post" summary:"任务提交"`
	Name   string
	Args   []Arg
}
type TaskSubmitRes struct {
	g.Meta `mime:"text/json"`
	TaskId string `json:"task_id"`
	Result any    `json:"result"`
}
