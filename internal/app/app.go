package app

import (
	"context"
	psql "melvad/internal/adapter/dbs/postgres"
	rstorage "melvad/internal/adapter/dbs/redis"
	"melvad/internal/adapter/hasher"
	"melvad/internal/config"
	"melvad/internal/service"
	"melvad/internal/transport/rest"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

func Run(cfg *config.Config) error {
	mux := http.NewServeMux()
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})
	pdb, err := pgx.Connect(context.Background(), cfg.Psql.Url)
	if err != nil {
		return err
	}
	_, err = pdb.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users(id serial PRIMARY KEY, name VARCHAR(100), age INT);")
	if err != nil {
		return err
	}

	redis := rstorage.New(rdb)
	store := psql.New(pdb)
	hasher := hasher.New()
	svc := service.New(redis, store, hasher)
	api := rest.New(svc)
	api.RegisterRoutes(mux)

	s := http.Server{
		Addr:    cfg.HttpServer.Port,
		Handler: mux}
	err = s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
