package service

type Service struct {
	pgDatabase pgDatabase
	cache      memcache
	subscriber natsSubscriber
}
