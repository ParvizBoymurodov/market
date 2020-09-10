package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/ParvizBoymurodov/market/cmd/app"
	"github.com/ParvizBoymurodov/market/pkg/models"
	"github.com/ParvizBoymurodov/market/pkg/services"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
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


var (
	host = flag.String("host", "", "Server host")
	port = flag.String("port", "", "Server port")
)
const envHost = "HOST"
const envPort = "PORT"

func fromFLagOrEnv(flag *string, envName string) (server string, ok bool){
	if *flag != ""{
		return *flag, true
	}
	return os.LookupEnv(envName)
}

func main() {

	flag.Parse()
	hostf, ok := fromFLagOrEnv(host, envHost)
	if !ok {
		hostf = *host
	}
	portf, ok := fromFLagOrEnv(port, envPort)
	if !ok {
		portf = *port
	}
	addr := net.JoinHostPort(hostf, portf)
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