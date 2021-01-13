// https://www.alexedwards.net/blog/serving-static-sites-with-go

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func main() {
	var port int = 3000
	var fileServer http.Handler = http.FileServer(http.Dir("./static"))

	// Handlers
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", serveTemplate)

	// Log
	log.Println("Listening on " + `:` + strconv.Itoa(port))

	// Serve
	log.Fatal(http.ListenAndServe(`:`+strconv.Itoa(port), nil))
}

func serveTemplate(writer http.ResponseWriter, reader *http.Request) {
	layoutPath := filepath.Join("templates", "indexLayout.tmpl")
	filePath := filepath.Join("templates", filepath.Clean(reader.URL.Path))

	fmt.Println(filepath.Clean(reader.URL.Path))

	tmpl, err := template.ParseFiles(layoutPath, filePath)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	log.Fatal(tmpl.ExecuteTemplate(writer, "indexLayout", nil))
}
