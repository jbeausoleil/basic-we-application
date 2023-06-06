package main

import (
	"fmt"
	"github.com/jbeausoleil/modern-web-apps/02-template-web-application/pkg/config"
	"github.com/jbeausoleil/modern-web-apps/02-template-web-application/pkg/handlers"
	"github.com/jbeausoleil/modern-web-apps/02-template-web-application/pkg/render"
	"log"
	"net/http"
)

const portNumber = 8080

func main() {
	var app config.AppConfig

	// Assign CreateTemplateCache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// Listen for a request sent by a web function
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("staring application on port %d", portNumber))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", portNumber), nil))
}
