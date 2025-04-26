package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhakimsaputra17/tickitz-api-go/config"
)

func NewPostgresPool(cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		 "postgres://%s:%s@%s:%s/%s",
		 cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	ctx, cancel :=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err:= pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// Test Connection 
	if err:= pool.Ping(ctx); err!= nil{
		return nil, err
	}
	return pool, nil

}