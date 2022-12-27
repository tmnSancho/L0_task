package main

import (
	"log"
	"order_service/internal/api"
	"order_service/internal/config"
	"order_service/internal/model"
	"order_service/internal/service"
	"order_service/internal/store"
	"order_service/pkg/nats"

	"github.com/nats-io/stan.go"
)

func main() {
	ch := make(chan model.Order, 5)

	cache := cache.NewCache()
	db, err := repo.NewPgRepo(config.Cfg.PgCfg)
	if err != nil {
		log.Fatalf("NewPgRepo: err %s", err)
	}

	store := store.New(db, cache)

	nc, err := stan.Connect(config.Cfg.NatsCfg.ClusterID, config.Cfg.NatsCfg.ClientID)
	if err != nil {
		log.Fatalf("stan.Connect: err %s", err)
	}
	defer nc.Close()

	sub, err := nats.NewSubscription(nc, config.Cfg.NatsCfg, ch)
	if err != nil {
		log.Fatalf("nats.NewSubscription: err %s", err)
	}
	defer sub.Unsubscribe()
	defer close(ch)

	_ = service.New(store, ch)

	api.StartService(store)
}
