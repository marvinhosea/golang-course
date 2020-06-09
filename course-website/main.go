package main

import (
	"fmt"
	"net/http"
)

func handleLandingPage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<h1>Hello There Marvin</h1>")
}

func main() {
	http.HandleFunc("/", handleLandingPage)
	http.ListenAndServe(":8000", nil)
}
