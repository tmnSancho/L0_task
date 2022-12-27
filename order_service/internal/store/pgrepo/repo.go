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

func (r *PgRepo) Set(order model.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		fmt.Printf("Set: can't start tx err %s", err)
		return err
	}
	if err := r.setOrder(ctx, order); err != nil {
		tx.Rollback(ctx)
		fmt.Printf("Set: setOrder err %s", err)
		return err
	}
	if err := r.setDelivery(ctx, order); err != nil {
		tx.Rollback(ctx)
		fmt.Printf("Set: setDelivery err %s", err)
		return err
	}
	if err := r.setPayment(ctx, order); err != nil {
		tx.Rollback(ctx)
		fmt.Printf("Set: setPayment err %s", err)
		return err
	}
	if err := r.setItems(ctx, order); err != nil {
		tx.Rollback(ctx)
		fmt.Printf("Set: setItems err %s", err)
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *PgRepo) setOrder(ctx context.Context, order model.Order) error {
	_, err := r.db.Exec(ctx, `INSERT INTO orders (order_uid, track_number, entry, delivery_id, payment_id, "+
	"locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) "+
	"VALUES (:order_uid, :track_number, :entry, :delivery_id, :payment_id, :locale, :internal_signature, "+
	":customer_id, :delivery_service, :shardkey, :sm_id, :date_created, :oof_shard)`, &order)

	return err
}

func (r *PgRepo) setDelivery(ctx context.Context, order model.Order) error {
	_, err := r.db.Exec(ctx, `INSERT INTO deliverys (order_uid, name, phone, zip, city, address, region, email) " +
	"VALUES (:order_uid, :name, :phone, :zip, :city, :address, :region, :email)`, &order.OrderUID, &order.Delivery)

	return err
}

func (r *PgRepo) setPayment(ctx context.Context, order model.Order) error {
	_, err := r.db.Exec(ctx, `INSERT INTO payments (order_uid, transaction, request_id, currency, provider, "+
	"amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES (:order_uid, :transaction, :request_id, "+
	":currency, :provider, :amount, :payment_dt, :bank, :delivery_cost, :goods_total, :custom_fee)`, &order.OrderUID, &order.Payment)

	return err
}

func (r *PgRepo) setItems(ctx context.Context, order model.Order) error {
	for _, item := range order.Items {
		_, err := r.db.Exec(ctx, `INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, "+
		"total_price, nm_id, brand, status, order_uid) VALUES (:chrt_id, :track_number, :price, :rid, :name, :sale, :size, "+
		":total_price, :nm_id, :brand, :status, :order_uid)`, &item)

		if err != nil {
			return err
		}
	}
	return nil
}
