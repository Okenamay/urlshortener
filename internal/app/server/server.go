package server

import (
	"net/http"

	"github.com/Okenamay/urlshortener/internal/app/configs"
	"github.com/Okenamay/urlshortener/internal/app/server/handlers"
)

func Launch() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.AutoHandler)

	serv := http.Server{
		Addr:        configs.ServerPort,
		Handler:     mux,
		IdleTimeout: configs.IdleTimeout,
	}

	err := serv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
