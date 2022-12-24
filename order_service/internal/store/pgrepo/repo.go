package repo

import (
	"context"
	"fmt"
	"time"

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

func NewPgRepo(cfg Config) *PgRepo {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		fmt.Errorf("NewPgOrderRepo-ParseConfig-err %s", err)
		return nil
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Errorf("NewPgRepo-ConnectConfig-err %s", err)
		return nil
	}

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = time.Second
	}

	return &PgRepo{
		db:      db,
		timeout: timeout,
	}
}
