package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "World!"
	}

	fmt.Fprintf(w, "Hello , %s", name)

}

func sayBye(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World!"
	}

	fmt.Fprintf(w, "Bye , %s", name)
}
