package main

import (
	"github.com/go-chi/chi"
	"morandev.com/views"
	"net/http"
)

var homePage *views.View
var contactPage *views.View
var faqPage *views.View

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	must(homePage.Render(writer, nil))
}

func contact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	must(contactPage.Render(writer, nil))
}

func faq(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	must(faqPage.Render(writer, nil))
}

func main() {
	contactPage = views.NewView("appLayout","views/contact.gohtml")
	homePage = views.NewView("appLayout","views/home.gohtml")
	faqPage = views.NewView("appLayout","views/faq.gohtml")
	
	router := chi.NewRouter()
	router.Get("/", home)
	router.Get("/contact", contact)
	router.Get("/faqs", faq)
	//router.NotFoundHandler = http.HandlerFunc(notFound)
	_ = http.ListenAndServe(":8000", router)
}

func must(error error)  {
	if error != nil {
		panic(error)
	}
}
