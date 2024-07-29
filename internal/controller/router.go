package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"task-scheduler/internal/controller/task"
)

func RegisterAPI(s *ghttp.Server) {

	s.BindMiddlewareDefault(ghttp.MiddlewareHandlerResponse)

	s.Group("/task", func(group *ghttp.RouterGroup) {
		group.Bind(
			task.NewV1(),
		)
	})
}
