// https://www.alexedwards.net/blog/serving-static-sites-with-go

package main

import(
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", serveTemplate)

	log.Println("Listening on :3000...")
	err:= http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "indexLayout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	fmt.Println(filepath.Clean(r.URL.Path))

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(w, "indexLayout", nil)
	if err != nil {
		log.Fatal(err)
	}
}