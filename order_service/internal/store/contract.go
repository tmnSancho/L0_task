package store

import "order_service/internal/model"

type pgRepo interface {
	GetDataForCache() ([]model.Order, error)
}

type memCache interface {
	GetOrderFromCache(orderUid string) *model.Order
	UploadCache(orders []model.Order) error
}
