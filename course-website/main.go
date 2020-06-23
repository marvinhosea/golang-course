package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homePage *template.Template
var contactPage *template.Template

func home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	err := homePage.Execute(writer, nil)
	if err != nil {
		panic(err)
	}
}

func contact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	err := contactPage.Execute(writer, nil)
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
	fmt.Fprint(writer, "This is a FAQ page for our good users <br /> <ul><li>What we do during our spare time</li><li>Most of the time we read novels and tell each other stories about star trek</li></ul>")
}

func main() {
	var err error

	homePage, err = template.ParseFiles("views/home.gohtml",
		"views/layouts/footer.gohtml")

	if err != nil {
		panic(err)
	}

	contactPage, err = template.ParseFiles("views/contact.gohtml",
		"views/layouts/footer.gohtml")

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":8000", router)
}
