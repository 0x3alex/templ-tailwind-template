package main

import (
	"embed"
	_ "embed"
	"github.com/0x3alex/templ-tailwind-template/internal/templates/index"
	"github.com/a-h/templ"
	"net/http"
)

//go:generate npm run build

//go:embed static
var static embed.FS

func main() {
	router := http.NewServeMux()
	router.Handle("/static/", http.FileServer(http.FS(static)))
	router.Handle("/", templ.Handler(index.Index("Hi")))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		println(err.Error())
	}
}
