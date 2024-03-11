package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/Hola", holaHandle)
	http.ListenAndServe(":8080", nil)

}

func holaHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hola go web")
}
