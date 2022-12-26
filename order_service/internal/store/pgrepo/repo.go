package repo

import (
	"context"
	"fmt"
	"order_service/internal/model"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PgRepo struct {
	db      *pgxpool.Pool
	timeout time.Duration
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	Timeout  time.Duration
}

func NewPgRepo(cfg Config) (*PgRepo, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("NewPgRepo: ConnectConfig err %s", err)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("NewPgRepo: ConnectConfig err %s", err)
	}

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = time.Second
	}

	return &PgRepo{
			db:      db,
			timeout: timeout,
		},
		nil
}

func (r *PgRepo) GetDataForCache() ([]model.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	orders := make([]model.Order, 0)

	if err := pgxscan.Get(ctx, r.db, &orders, `SELECT * FROM orders`); err != nil {
		return nil, fmt.Errorf("GetDataForCache: Get err %s", err)
	}

	for _, order := range orders {
		if err := pgxscan.Get(ctx, r.db, &order.Delivery, `SELECT * FROM deliverys WHERE id = $1`, order.OrderUID); err != nil {
			return nil, fmt.Errorf("GetDataForCache: Get deliverys err %s", err)
		}

		if err := pgxscan.Get(ctx, r.db, &order.Payment, `SELECT * FROM payments WHERE id = $1`, order.OrderUID); err != nil {
			return nil, fmt.Errorf("GetDataForCache: Get deliverys err %s", err)
		}

		if err := pgxscan.Get(ctx, r.db, &order.Items, `SELECT * FROM items WHERE id = $1`, order.OrderUID); err != nil {
			return nil, fmt.Errorf("GetDataForCache: Get deliverys err %s", err)
		}
	}

	return orders, nil
}
