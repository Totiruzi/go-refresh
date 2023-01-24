package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Totiruzi/go-refresh/pkg/config"
	"github.com/Totiruzi/go-refresh/pkg/handlers"
	"github.com/Totiruzi/go-refresh/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCatche()
	if err != nil {
		log.Fatal("could not create template cache")
	}
	app.TamplateCatch = tc
	app.UseCatche = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("server is starting on port", port)
	http.ListenAndServe(port, nil)
}
