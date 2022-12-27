package service

import (
	"log"
	"order_service/internal/model"
)

type Service struct {
	store store
	ch    chan model.Order
}

func New(store store, ch chan model.Order) *Service {
	s := &Service{store: store, ch: ch}
	go s.ReciveFromChan()
	return s
}

func (s *Service) ReciveFromChan() {
	if err := s.store.UploadCache(); err != nil {
		log.Fatalf("ReciveFromChan: err %s", err)
	}

	for {
		order := <-s.ch

		if err := s.store.Set(order); err != nil {
			log.Printf("ReciveFromChan: Set cannot insert order %s into store err: %s", order.OrderUID, err)
		}
	}
}
