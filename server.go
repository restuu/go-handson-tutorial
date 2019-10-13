package main

import (
	"fmt"
	"log"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %q", r.URL.Path)

}

func main() {
	http.HandleFunc("/", get)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
