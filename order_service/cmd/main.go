package main

import (
	"fmt"
	"log"
	"order_service/internal/model"
)

func main() {
	log.Fatal()
	m := model.Order{}
	fmt.Print(m)

	db := repo.NewPgOrderRepo()
}
