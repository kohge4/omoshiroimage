package main

import (
	"context"
	"log"
	"time"

	"omoshiroimg/external"
	"omoshiroimg/web"
)

func main() {

	srv := web.NewServer()
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	//GenerateImage()
	imgen := external.NewChromedpImageGenerator()
	imgen.GenerateImage()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Print(err)
	}
	return
}
