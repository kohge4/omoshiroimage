package web

import (
	"fmt"
	"log"
	"net/http"
	"omoshiroimg/external"
)

func (app *Handler) ImageGenerator(w http.ResponseWriter, r *http.Request) {
	//image, ok := r.Context().Value("img-path").(string)
	//if !ok {
	//	image = ""
	//}
	var image string
	if r.FormValue("image") != "" {
		image = fmt.Sprintf("/tmp/%v.png", r.FormValue("image"))
	}
	fmt.Println(image)

	if err := app.View["image_generator"].Execute(w, struct {
		ImageURL string
	}{
		ImageURL: image,
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func (app *Handler) ImageGeneratorExec(w http.ResponseWriter, r *http.Request) {

	// https://pkg.go.dev/net/http#Request.FormValue
	message := r.FormValue("message")
	fmt.Println(message)
	imgen := external.NewChromedpImageGenerator()
	imgPath := imgen.GenerateImage(message)

	// TODO リダイレクトして画像保存
	//ctx := context.WithValue(r.Context(), "img-path", imgPath)
	//r = r.WithContext(ctx)
	//image, _ := r.Context().Value("img-path").(string)
	//fmt.Println("context", image)

	http.Redirect(w, r, fmt.Sprintf("/imgen?image=%v", imgPath), http.StatusFound)
}

type bodyJSON struct {
	Message string `json:"message"`
}
