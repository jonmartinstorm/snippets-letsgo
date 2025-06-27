package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("starter server på", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		logger.Error("feil i starting av server", "error", err)
		os.Exit(1)
	}
}
