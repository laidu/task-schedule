package cmd

import (
	"context"
	"encoding/json"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"task-scheduler/internal/worker"
)

var (
	MachineryServer = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := NewMachineryServer(ctx)
			return worker.RegisterWorker(s)
		},
	}
)

func NewMachineryServer(ctx context.Context) *machinery.Server {

	cfg := g.Cfg("machinery")

	g.Redis("")

	configData, _ := cfg.Data(ctx)
	configJson, _ := json.Marshal(configData)

	cnf := &config.Config{}
	_ = json.Unmarshal(configJson, cnf)

	// Create server instance
	broker := redisbroker.NewGR(cnf, []string{"localhost:6379"}, 0)
	backend := redisbackend.NewGR(cnf, []string{"localhost:6379"}, 0)
	lock := eagerlock.New()

	return machinery.NewServer(cnf, broker, backend, lock)
}
