package store

import "order_service/internal/model"

type pgRepo interface {
	GetData() []model.Order
}

type memCache interface {
}
