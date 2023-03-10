package cache

import (
	"order_service/internal/model"
	"sync"
)

type Cache struct {
	cache map[string]model.Order
	mutex *sync.RWMutex
}

func NewCache() *Cache {
	m := make(map[string]model.Order)
	mut := &sync.RWMutex{}

	return &Cache{
		cache: m,
		mutex: mut,
	}
}

func (c *Cache) GetOrderFromCache(orderUid string) *model.Order {
	c.mutex.RLock()
	order, ok := c.cache[orderUid]
	if !ok {
		return nil
	}
	c.mutex.RUnlock()

	return &order
}

func (c *Cache) UploadCache(orders []model.Order) error {
	for _, order := range orders {
		c.mutex.Lock()
		c.cache[order.OrderUID] = order
		c.mutex.Unlock()
	}

	return nil
}

func (c *Cache) GetOrderById(id string) model.Order {
	c.mutex.RLock()
	order := c.cache[id]
	c.mutex.RUnlock()
	return order
}

func (c *Cache) Set(order model.Order) {
	c.mutex.RLock()
	c.cache[order.OrderUID] = order
	c.mutex.RUnlock()
}
