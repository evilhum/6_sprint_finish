package main

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	srv, err := server.NewServer(logger)
	if err != nil {
		logger.Fatal("failed to create server: ", err)
	}

	logger.Println("starting server on :8080")
	if err := http.ListenAndServe(srv.Server.Addr, srv.Server.Handler); err != http.ErrServerClosed {
		logger.Fatal("server failed: ", err)
	}
}
