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
	var port = 3000
	var fileServer = http.FileServer(http.Dir("./static"))

	// Handlers
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", serveTemplate)

	// Log
	log.Println("Listening on " + `:` + strconv.Itoa(port))

	// Serve
	log.Fatal(http.ListenAndServe(`:`+strconv.Itoa(port), nil))
}

func serveTemplate(writer http.ResponseWriter, reader *http.Request) {
	// Get Paths
	var layoutPath = filepath.Join("templates", "indexLayout.tmpl")
	var filePath = filepath.Join("templates", filepath.Clean(reader.URL.Path))
	fmt.Println("Serving Path: " + filepath.Clean(reader.URL.Path))

	// Parse Template
	tmpl, err := template.ParseFiles(layoutPath, filePath)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// Execute Template
	log.Fatal(tmpl.ExecuteTemplate(writer, "indexLayout", nil))
}
