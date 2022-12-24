package service

type Service struct {
	pgDatabase pgDatabase
	cache      memCache
	subscriber natsSubscriber
}
