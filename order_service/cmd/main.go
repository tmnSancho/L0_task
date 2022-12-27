package main

import (
	"fmt"
	"log"
	"order_service/internal/model"
	"order_service/pkg/nats"
)

type Config struct {
	natsCfg nats.Config
}

func main() {
	log.Fatal()
	m := model.Order{}
	fmt.Print(m)

	db := repo.NewPgRepo()
}
