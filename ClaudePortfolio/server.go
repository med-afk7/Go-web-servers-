package main

import (
	"log"
	"net/http"
)

func servePage(file string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/pages/"+file)
	}
}

func main() {
	// CSS static files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))

	// "/" must be exact — otherwise it swallows every unknown route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "static/pages/hero.html")
	})

	http.HandleFunc("/about", servePage("/about.html"))
	http.HandleFunc("/experience", servePage("/experience.html"))
	http.HandleFunc("/projects", servePage("/projects.html"))
	http.HandleFunc("/skills", servePage("/skills.html"))
	http.HandleFunc("/contact", servePage("/contact.html"))

	log.Println("Running → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
