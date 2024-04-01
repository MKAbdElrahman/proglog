package main

import (
	"log"
	"proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":3000")
	log.Fatal(srv.ListenAndServe())
}
