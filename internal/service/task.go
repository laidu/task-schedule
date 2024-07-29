package service

import (
	"context"
	"encoding/json"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/gogf/gf/v2/frame/g"
	"sync"
)

var once sync.Once

var machineryServer *machinery.Server

func NewMachineryServer() *machinery.Server {

	once.Do(func() {
		cfg := g.Cfg("machinery")

		configData, _ := cfg.Data(context.Background())
		configJson, _ := json.Marshal(configData)

		cnf := &config.Config{}
		_ = json.Unmarshal(configJson, cnf)

		// Create server instance
		broker := redisbroker.NewGR(cnf, []string{"10.0.0.110:6379"}, 0)
		backend := redisbackend.NewGR(cnf, []string{"10.0.0.110:6379"}, 0)
		lock := eagerlock.New()
		machineryServer = machinery.NewServer(cnf, broker, backend, lock)
	})

	return machineryServer
}
