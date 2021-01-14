// Quick semi-static golang server
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var port = 3000
	var fileServer = http.FileServer(http.Dir("./static"))

	// Handlers
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", serveTemplate)

	// Log
	log.Println("Listening on " + `:` + strconv.FormatInt(int64(port), 10))

	// Serve
	log.Fatal(http.ListenAndServe(`:`+strconv.FormatInt(int64(port), 10), nil))
}

func serveTemplate(writer http.ResponseWriter, reader *http.Request) {
	// Get Paths
	var layoutPath = filepath.Join("templates", strings.Split(reader.URL.Path, ".")[0]+"Layout.tmpl")
	var filePath = filepath.Clean("./" + reader.URL.Path)

	fmt.Println("Serving Path: " + reader.URL.Path)

	// Return if 404
	is404 := handle404(writer, filePath)
	if is404 {
		fmt.Println("Server returned 404.")
		return
	}

	// Parse Template
	tmpl, err := template.ParseFiles(layoutPath, filePath)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(writer, http.StatusText(500), 500)
		return
	}

	// Execute Template
	err = tmpl.ExecuteTemplate(writer, "indexLayout", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, http.StatusText(500), 500)
	}
}

func handle404(writer http.ResponseWriter, filePath string) bool {
	// Return a 404 if the template doesn't exist
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			template404(writer)
			return true
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		template404(writer)
		return true
	}

	return false
}

func template404(writer http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/404Layout.tmpl", "404.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(writer, "404Layout", nil)
	if err != nil {
		log.Fatal(err)
	}
}
