package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/stefanovazzocell/TuringMachine/src/api"
)

var (
	logLevel slog.Level

	serverAddr  string
	corsOrigins string

	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration

	gamesDbFile    string
	dbForceRefresh bool
)

func init() {
	flag.StringVar(&serverAddr, "addr", "0.0.0.0:8080", "the address to listen from")
	flag.StringVar(&corsOrigins, "cors_origin", "*", "the CORS origin")

	flag.DurationVar(&readTimeout, "read_timeout", 5*time.Second, "timeout for the request read")
	flag.DurationVar(&writeTimeout, "write_timeout", 10*time.Second, "timeout for the request read+write")
	flag.DurationVar(&idleTimeout, "idle_timeout", 2*time.Minute, "the keepalive timeout between requests")

	flag.StringVar(&gamesDbFile, "db", "./games", "the location of the games DB file")
	flag.BoolVar(&dbForceRefresh, "db_force_refresh", false, "if set, forces the database refresh at startup")

	flag.TextVar(&logLevel, "log_level", slog.LevelInfo, "sets the log level")

	flag.Parse()

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	})))
}

func main() {
	config := api.NewAPIConfig(gamesDbFile, corsOrigins)
	config.StoreForceCreate = dbForceRefresh

	a, err := api.NewApi(&http.Server{
		Addr: serverAddr,

		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}, config)
	if err != nil {
		panic(err)
	}
	defer a.Close()

	a.ListenAndServe()
	a.AwaitInterrupt()
}
