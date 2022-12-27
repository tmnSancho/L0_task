package service

import "order_service/internal/model"

type store interface {
	GetOrderById(id string) *model.Order
	UploadCache() error
}
