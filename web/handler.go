package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	View map[string]*template.Template
	//Repository domain.ArticleRepository
}

func NewHandler(template map[string]*template.Template) *Handler {
	return &Handler{
		//DB:       db,
		View: template,
	}
}

func (app *Handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	url := "assets/image/glassp.png"
	text := "おはようございます"

	if err := app.View["card"].Execute(w, struct {
		ImageURL string
		Text     string
	}{
		ImageURL: url,
		Text:     text,
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func (app *Handler) SelfIntroduction(w http.ResponseWriter, r *http.Request) {
	url := "assets/image/glassp.png"
	name := "koge"
	title := "テスト奴"
	text := "おはようございます"

	if err := app.View["selfintroduction"].Execute(w, struct {
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

func (app *Handler) Fukidashi(w http.ResponseWriter, r *http.Request) {
	url := "assets/image/glassp.png"
	name := "koge"
	title := "テスト奴"
	text := "おはようございます"

	if err := app.View["fukidashi"].Execute(w, struct {
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

func responseByJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
	return
}
