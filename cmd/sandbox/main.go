package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title    string
	Subtitle string
}

func main() {
	tmpl := template.Must(template.ParseFiles("web/templates/base.tmpl"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:    "Form 1040: High Net Worth 2026",
			Subtitle: "David Vance · 6+ Years Tax Technology Cohort",
		}
		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("template error: %v", err)
			http.Error(w, "template error", http.StatusInternalServerError)
		}
	})

	log.Println("ProConnect Tax Engine (Premium Design) running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
