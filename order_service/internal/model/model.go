package model

import "time"

type Order struct {
	OrderUID          string    `json:"order_uid" db:"order_uid"`
	TrackNumber       string    `json:"track_number" db:"track_number"`
	Entry             string    `json:"entry" db:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerId        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	SmId              int64     `json:"sm_id" db:"sm_id"`
	DateCreated       string    `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
	LastAccessTime    time.Time `json:"-"` //for deleting order from cache
}

type Delivery struct {
	OrderUID string `json:"-" db:"order_uid"`
	Name     string `json:"name" db:"name"`
	Phone    string `json:"phone" db:"phone"`
	Zip      string `json:"zip" db:"zip"`
	City     string `json:"city" db:"city"`
	Address  string `json:"address" db:"address"`
	Region   string `json:"region" db:"region"`
	Email    string `json:"email" db:"email"`
}

type Payment struct {
	OrderUID     string `json:"-" db:"order_uid"`
	Transaction  string `json:"transaction" db:"transaction"`
	RequestId    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int64  `json:"amount" db:"amount"`
	PaymentDt    int64  `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int64  `json:"delivery_cost" db:"delivery_cost"`
	CustomFee    string `json:"custom_fee" db:"custom_fee"`
}

type Item struct {
	OrderUID    string `json:"-" db:"order_uid"`
	ChrtId      int64  `json:"chrt_id" db:"chrt_id"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Price       int64  `json:"price" db:"price"`
	Rid         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int64  `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  string `json:"total_price" db:"total_price"`
	NmId        string `json:"nm_id" db:"nm_id"`
	Brand       string `json:"brand" db:"brand"`
	Status      string `json:"status" db:"status"`
}
