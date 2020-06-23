package main

import (
	"fmt"
	"morandev.com/views"
	"net/http"

	"github.com/gorilla/mux"
)

var homePage *views.View
var contactPage *views.View
var faqPage *views.View

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	err := homePage.Template.Execute(writer, nil)
	if err != nil {
		panic(err)
	}
}

func contact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	err := contactPage.Template.Execute(writer, nil)
	if err != nil {
		panic(err)
	}
}

func notFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprint(writer, "The page you are looking is not found")
}

func faq(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	err := faqPage.Template.Execute(writer, nil)

	if err != nil {
		panic(err)
	}
}

func main() {
	homePage = views.NewView("views/home.gohtml")
	contactPage = views.NewView("views/contact.gohtml")
	faqPage = views.NewView("views/faq.gohtml")
	
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":8000", router)
}
