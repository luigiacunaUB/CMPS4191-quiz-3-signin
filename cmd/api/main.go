// Luigi Acuna
// CMPS4191 Quiz 3
// main.go 1
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const appVersion = "1.0.0"

type serverConfig struct {
	port       int    //port number to access signin page
	enviroment string //enviroment the signin page will be on
	db         struct {
		dsn string
	}
}

type applicationDependencies struct {
	config serverConfig
	logger *slog.Logger //look more into this later
}

func main() {
	var settings serverConfig

	//Settings ports and enviroment info
	flag.IntVar(&settings.port, "port", 4000, "Server Port")
	flag.StringVar(&settings.enviroment, "env", "development", "Enviroment(development|staging|)")
	flag.StringVar(&settings.db.dsn, "db-dsn", "postgres://signin:password123@localhost/signin?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Open connection
	db, err := openDB(settings)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	//release db resources before closing
	defer db.Close()
	logger.Info("database connection before exiting")

	appInstance := &applicationDependencies{
		config: settings,
		logger: logger,
	}

	//api server info
	apiServer := http.Server{
		Addr:         fmt.Sprintf(":%d", settings.port),
		Handler:      appInstance.routes(), //this one too
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "address", apiServer.Addr, "enviroment", settings.enviroment)

	err = apiServer.ListenAndServe()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(settings serverConfig) (*sql.DB, error) {
	//open a connection pool
	db, err := sql.Open("postgres", settings.db.dsn)
	if err != nil {
		return nil, err
	}

	//set context to ensure DB ops dont take too long
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//test
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
