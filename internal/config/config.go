package config

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrArgs = errors.New("wrong args")
)

type Config struct {
	Redis      *Redis
	Psql       *Postgres
	HttpServer *HttpServer
}
type Redis struct {
	Addr string
}
type Postgres struct {
	Url string
}
type HttpServer struct {
	Port string
}

func New() (*Config, error) {
	args := os.Args
	connAddr, err := validateArgs(args)
	if err != nil {
		return &Config{}, err
	}
	return &Config{
		Redis: &Redis{
			Addr: connAddr,
		},
		Psql: &Postgres{
			Url: "postgres://postgres:Wild54323@localhost:5432/postgres",
		},
		HttpServer: &HttpServer{
			Port: "localhost:8080",
		},
	}, nil
}
func validateArgs(args []string) (string, error) {
	if len(args) != 5 {
		return "", ErrArgs
	}
	if args[1] != "-host" || args[3] != "-port" {
		return "", ErrArgs
	}
	return fmt.Sprintf("%s:%s", args[2], args[4]), nil
}
