package store

import "order_service/internal/model"

type Store struct {
	db    pgRepo
	cache memCache
}

func New(db pgRepo, cache memCache) *Store {
	return &Store{db: db, cache: cache}
}

func (s *Store) UploadCache() error {
	data, err := s.db.GetDataForCache()
	if err != nil {
		return err
	}

	if err := s.cache.UploadCache(data); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetOrderById(id string) *model.Order {
	return s.cache.GetOrderFromCache(id)
}
