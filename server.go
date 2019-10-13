package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %q", r.URL.Path)
}

func getJSON(w http.ResponseWriter, r *http.Request) {
	data := struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		MailAddress string `json:"mail_address"`
	}{
		ID:          1,
		Name:        "John",
		MailAddress: "John@mail.com",
	}

	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Write(byte)

}

func handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not found, %q", r.URL.Path)
}

func main() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/json", getJSON)

	http.HandleFunc("/404", handle404)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
