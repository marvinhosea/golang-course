package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<h1>Hello There Marvin</h1>")
}

func contact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<h1>This is a contact page</h1> <br /> Email us: marvin@appslab.co.ke")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	http.ListenAndServe(":8000", router)
}
