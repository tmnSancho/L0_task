package service

type pgDatabase interface {
	SaveOrder()
}

type memcache interface {
	GetCacheFromDB()
}

type natsSubscriber interface {
}
