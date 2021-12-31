package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"omoshiroimg/web"
)

func main() {

	var templates = make(map[string]*template.Template)
	templates["card"] = loadTemplate("card")
	templates["selfintroduction"] = loadTemplate("selfintroduction")
	//infra.NewElasticSearchClient()
	handler := web.NewHandler(templates)

	r := mux.NewRouter().StrictSlash(true)
	r.Use(loggingMiddleware)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	r.HandleFunc("/", handler.IndexPage)
	r.HandleFunc("/me", handler.SelfIntroduction)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	return
}

func htmlHandler(name string) func(w http.ResponseWriter, r *http.Request) {
	temp := fmt.Sprintf("assets/template/%s.html", name)
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(
			temp,
		)
		if err != nil {
			log.Fatalf("template error: %v", err)
		}
		if err := t.Execute(w, struct {
			Text string
		}{
			Text: "hoge",
		}); err != nil {
			log.Printf("failed to execute template: %v", err)
		}
	}
}

func loadTemplate(name string) *template.Template {
	t, err := template.ParseFiles(
		"assets/template/"+name+".html",
		"assets/template/_header.html",
		"assets/template/_footer.html",
	)
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	return t
}

// 全部のファイルを読み込む奴
func localAlltemplate() {}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		t := t2.Sub(t1)
		log.Printf("[%s] %s %s", r.Method, r.URL, t.String())
	})
}
