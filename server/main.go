package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rinosukmandityo/google-search-ext/server/api"
)

func main() {
	port := 8080

	mux := http.DefaultServeMux

	mux.HandleFunc("/search", api.PostSearch)

	var handler http.Handler = mux
	handler = api.CorsMiddleware(handler)

	server := new(http.Server)
	server.Addr = fmt.Sprintf(":%d", port)
	server.Handler = handler

	log.Printf("server started at localhost:%d\n", port)
	server.ListenAndServe()
}
