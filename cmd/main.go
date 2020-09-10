package main

import (
	"context"
	"encoding/json"
	"github.com/ParvizBoymurodov/market/cmd/app"
	"github.com/ParvizBoymurodov/market/pkg/models"
	"github.com/ParvizBoymurodov/market/pkg/services"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"path/filepath"
)

var conf models.Config

func init() {
	data, err := ioutil.ReadFile("./cmd/configs.json")
	if err != nil {
		log.Fatalf("Can't read file : %v\n", err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("Can't unmarshal data : %v\n", err)
	}
}

func main() {
	addr := net.JoinHostPort(conf.Host, conf.Port)
	start(addr, conf.Dsn)
}

func start(addr string,dsn string) {
	router := app.NewExactMux()

	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	nellyMarket := services.NewNellyMarket(pool)

	server := app.NewServer(
		pool,
		router,
		nellyMarket,
		filepath.Join("web", "templates"),
		filepath.Join("web", "assets"),
	)

	server.InitRoutes()

	nellyMarket.Start()

	panic(http.ListenAndServe(addr,server))
}