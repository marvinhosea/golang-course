package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"morandev.com/views"
	"net/http"
)

var homePage *views.View
var contactPage *views.View
var faqPage *views.View

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	err := homePage.Template.ExecuteTemplate(writer, homePage.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	err := contactPage.Template.ExecuteTemplate(writer, contactPage.Layout, nil)
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
	err := faqPage.Template.ExecuteTemplate(writer, faqPage.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	contactPage = views.NewView("appLayout","views/contact.gohtml")
	homePage = views.NewView("appLayout","views/home.gohtml")
	faqPage = views.NewView("appLayout","views/faq.gohtml")
	
	router := chi.NewRouter()
	router.Get("/", home)
	router.Get("/contact", contact)
	router.Get("/faq", faq)
	//router.NotFoundHandler = http.HandlerFunc(notFound)
	_ = http.ListenAndServe(":8000", router)
}
