package service

type store struct {
}

type natsSubscriber interface {
	Subscribe()
}
