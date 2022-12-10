package app

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/CleysonPH/reading-tracker/config"
	"github.com/CleysonPH/reading-tracker/internal/database"
	"github.com/CleysonPH/reading-tracker/internal/repository"
	"github.com/CleysonPH/reading-tracker/internal/service"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/handler"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/middleware"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/validator"
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

	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logErr := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	loggerService := service.NewLoggerService(logInfo, logErr)
	loggerMiddleware := middleware.NewLoggerMiddleware(loggerService)

	loggerService.Info("Starting server on %s in %s mode", config.Addr(), config.Env)

	bookRepository := repository.NewBookModel(db)
	bookValidator := validator.NewBookValidator(bookRepository)
	bookHandler := handler.NewBookHandler(bookRepository, bookValidator)

	readingSessionRepository := repository.NewReadingSessionModel(db)
	readingSessionValidator := validator.NewReadingSessionValidator(bookRepository, readingSessionRepository)
	readingSessionHandler := handler.NewReadingSessionHandler(bookRepository, readingSessionValidator, readingSessionRepository)

	router := rest.NewRouter(bookHandler, readingSessionHandler)
	router = loggerMiddleware.Use(router)

	srv := &http.Server{
		Addr:    config.Addr(),
		Handler: router,
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
