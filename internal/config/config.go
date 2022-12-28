package config

import (
	"order_service/pkg/nats"
	"order_service/internal/pgrepo"
)

type Config struct {
	NatsCfg nats.Config
	PgCfg   repo.Config
}

var Cfg Config = Config{
	NatsCfg: nats.Config{
		ClusterID: "test-cluster",
		ClientID:  "reader",
		Channel:   "session",
	},
	PgCfg: repo.Config{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DbName:   "orderDB",
	},
}
