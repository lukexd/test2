package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/lubiePlacki", secPage)
	fmt.Println("Listening on port 8000")

	http.ListenAndServe(":8000", mux)
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home</h1>"))
}

func secPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Lubię Placki</h1>"))

}
