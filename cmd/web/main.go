package main

import (
	"log"

	"omoshiroimg/web"
)

func main() {
	srv := web.NewServer()
	log.Fatal(srv.ListenAndServe())

	return
}
