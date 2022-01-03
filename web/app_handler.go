package web

import (
	"fmt"
	"log"
	"net/http"
	"omoshiroimg/external"
)

func (app *Handler) ImageGenerator(w http.ResponseWriter, r *http.Request) {
	url := "assets/image/glassp.png"
	name := "koge"
	title := "テスト奴"
	text := "おはようございます"

	if err := app.View["image_generator"].Execute(w, struct {
		ImageURL string
		Title    string
		Text     string
		Name     string
	}{
		ImageURL: url,
		Title:    title,
		Text:     text,
		Name:     name,
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func (app *Handler) ImageGeneratorExec(w http.ResponseWriter, r *http.Request) {
	url := "assets/image/glassp.png"
	name := "koge"
	title := "テスト奴"
	text := "おはようございます"

	// https://pkg.go.dev/net/http#Request.FormValue
	message := r.FormValue("message")
	fmt.Println(message)
	imgen := external.NewChromedpImageGenerator()
	imgen.GenerateImage()

	// TODO リダイレクトして画像保存

	if err := app.View["image_generator"].Execute(w, struct {
		ImageURL string
		Title    string
		Text     string
		Name     string
	}{
		ImageURL: url,
		Title:    title,
		Text:     text,
		Name:     name,
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

type bodyJSON struct {
	Message string `json:"message"`
}
