package app

import (
	"flag"
	"log"
	"net/http"

	"github.com/CleysonPH/reading-tracker/config"
	"github.com/CleysonPH/reading-tracker/internal/database"
	"github.com/CleysonPH/reading-tracker/internal/repository"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/handler"
)

func Run() {
	flag.StringVar(&config.Port, "port", config.Port, "Port to run the server on")
	flag.StringVar(&config.Env, "env", config.Env, "Environment to run the server on")
	flag.StringVar(&config.Dsn, "dsn", config.Dsn, "Data source name to connect to the database")
	flag.StringVar(&config.Host, "host", config.Host, "Host to run the server on")
	flag.Parse()

	database.InitMysql(config.Dsn)
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookRepository := repository.NewBookModel(db)
	bookHandler := handler.NewBookHandler(bookRepository)

	router := rest.NewRouter(bookHandler)

	srv := &http.Server{
		Addr:    config.Addr(),
		Handler: router,
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
