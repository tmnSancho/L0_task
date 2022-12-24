package cache

import (
	"order_service/internal/model"
	"sync"
)

type Cache struct {
	m       map[string]model.Order
	maxSize int
	sync.RWMutex
}

func NewCache(cacheSize int) *Cache {
	m := make(map[string]model.Order, cacheSize)
	return &Cache{
		m:       m,
		maxSize: cacheSize,
	}
}

func (c *Cache) GetOrder(orderUid string) *model.Order {
	order, ok := c.m[orderUid]
	if !ok {
		return nil
	}

	return &order
}
