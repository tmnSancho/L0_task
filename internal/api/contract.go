package api

import "order_service/internal/model"

type store interface {
	UploadCache() error
	GetOrderById(id string) *model.Order
	Set(order model.Order) error
}
