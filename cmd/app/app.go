package app

import (
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"nellymarket/pkg/services"
	"net/http"
)

type server struct {
	pool          *pgxpool.Pool
	router        http.Handler
	marketSvc     *services.NellyMarket
	templatesPath string
	assetsPath    string
}

func NewServer(pool *pgxpool.Pool, router http.Handler, marketSvc *services.NellyMarket, templatesPath string, assetsPath string) *server {
	if pool == nil {
		panic(errors.New("Pool can't be nil"))
	}
	if router == nil {
		panic(errors.New("Router can't be nil"))
	}
	if marketSvc == nil {
		panic(errors.New("marketSvc can't be nil"))
	}
	if templatesPath == "" {
		panic(errors.New("templatesPath can't be empty"))
	}
	if assetsPath == "" {
		panic(errors.New("assetsPath can't be empty"))
	}
	return &server{
		pool:          pool,
		router:        router,
		marketSvc:     marketSvc,
		templatesPath: templatesPath,
		assetsPath:    assetsPath,
	}
}

func (receiver *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	receiver.router.ServeHTTP(w, r)
}
