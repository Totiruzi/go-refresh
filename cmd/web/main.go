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

	fmt.Println("server is starting on port", port)

	srv := http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
