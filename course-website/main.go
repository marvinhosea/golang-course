package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleLandingPage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<h1>Hello There Marvin</h1>")
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", handleLandingPage)
	mux.HandleFunc("/demo", handleLandingPage)
	http.ListenAndServe(":8000", mux)
}
