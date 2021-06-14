package main

import (
	"log"

	"github.com/pomyslowynick/proglog/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
