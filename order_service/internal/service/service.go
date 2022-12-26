package service

type Service struct {
	store      store
	subscriber natsSubscriber
}
