package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		templates := []string{
			filepath.Join("web", "templates", "base.html"),
			filepath.Join("web", "templates", "navbar.html"),
			filepath.Join("web", "templates", "home.html"),
		}

		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			log.Printf("Error parsing templates: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "base.html", map[string]interface{}{
			"Title": "Home Page",
		})
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	// Handler for JSON data generation
	http.HandleFunc("/generate-json", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Ensure form data is parsed
		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form data: %v", err)
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		metadata := map[string]interface{}{
			"DCMI_Metadata": map[string]interface{}{
				"Title":       r.FormValue("title"),
				"Creator":     r.FormValue("creator"),
				"Subject":     r.FormValue("subject"),
				"Description": r.FormValue("description"),
				"Publisher":   r.FormValue("publisher"),
				// Assume Contributors is submitted as a comma-separated list
				"Contributors":               splitContributors(r.FormValue("contributors")),
				"Date":                       r.FormValue("date"),
				"Type":                       r.FormValue("type"),
				"Format":                     r.FormValue("format"),
				"Identifier":                 r.FormValue("identifier"),
				"Source":                     r.FormValue("source"),
				"Language":                   r.FormValue("language"),
				"Relation":                   r.FormValue("relation"),
				"Coverage":                   r.FormValue("coverage"),
				"Rights":                     r.FormValue("rights"),
				"DatasetDivisionDescription": r.FormValue("datasetDivisionDescription"),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(metadata); err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, "Error generating JSON", http.StatusInternalServerError)
		}
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// Helper function to split the contributors string into a slice
func splitContributors(contributors string) []string {
	if contributors == "" {
		return []string{}
	}
	return strings.Split(contributors, ",")
}
