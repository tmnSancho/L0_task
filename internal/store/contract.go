package store

import "order_service/internal/model"

type pgRepo interface {
	Set(order model.Order) error
	GetDataForCache() ([]model.Order, error)
}

type memCache interface {
	GetOrderFromCache(orderUid string) *model.Order
	UploadCache(orders []model.Order) error
	Set(order model.Order)
}
