package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wickywaa/gocourse1/pkg/config"
	"github.com/wickywaa/gocourse1/pkg/handlers"
	"github.com/wickywaa/gocourse1/pkg/render"
)

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	const port string = ":8080"

	fmt.Println(fmt.Sprintf("listening on port %s", port))

	srv := &http.Server{

		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
