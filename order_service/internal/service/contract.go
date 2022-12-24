package service

type pgDatabase interface {
	SaveOrder()
	GetOrder()
}

type memCache interface {
	UploadCache()
}

type natsSubscriber interface {
	ReciveMsg()
}
