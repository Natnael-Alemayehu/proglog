package main

import (
	"log"

	"github.com/natnael-alemayehu/proglog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
