package main

import (
	"fmt"
	"net/http"

	"github.com/Snehadas2005/REST-API-GO/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to API wordld"))
	})

	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Println("SERVER STARTING ON:", cfg.HTTPServer.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	fmt.Println("server started")

}
