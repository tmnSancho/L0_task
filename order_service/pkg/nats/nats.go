package nats

import (
	"encoding/json"
	"log"
	"order_service/internal/model"

	"github.com/nats-io/stan.go"
)

var ch chan model.Order

type Config struct {
	ClusterID string
	ClientID  string
	Channel   string
	URL       string
}

func handle(msg *stan.Msg) {
	var data model.Order

	err := json.Unmarshal(msg.Data, &data)
	if err != nil {
		log.Printf("error while decoding data from nats channel: %v ", err)
		msg.Ack()
		return
	}

	ch <- data

	if err := msg.Ack(); err != nil {
		log.Printf("failed tp ACK msg: %d", msg.Sequence)
		return
	}
}

func NewSubscription(conn stan.Conn, cfg Config, c chan model.Order) (stan.Subscription, error) {
	ch = c
	sub, err := conn.Subscribe(
		cfg.Channel,
		handle,
		stan.SetManualAckMode(),
	)

	if err != nil {
		return nil, err
	}
	return sub, nil
}
