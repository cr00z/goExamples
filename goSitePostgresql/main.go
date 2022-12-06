package main

import (
	"context"
	"github.com/cr00z/goSite/postgresqlSite/internal/application"
	"github.com/cr00z/goSite/postgresqlSite/internal/repository"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	dbpool, err := repository.InitDBConn(ctx)
	if err != nil {
		log.Fatalf("%w failed to init DB connection", err)
	}
	defer dbpool.Close()
	a := application.NewApp(ctx, dbpool)
	r := httprouter.New()
	a.Routes(r)
	srv := &http.Server{Addr: "0.0.0.0:8080", Handler: r}
	log.Println("It's alive! Try http://localhost:8080")
	srv.ListenAndServe()
}
