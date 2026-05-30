package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandleFunc)
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/bye", sayBye)
	http.ListenAndServe(":8080", nil)

}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	a := 10
	b := 20
	fmt.Fprintf(w, "Hello , World! %v", a+b)
}
