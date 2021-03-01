package main

import (
	"fmt"
	"net/http"

	"github.com/IAPOLINARIO/go-stuff/web-app/views"
	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

func Handle404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "Sorry your page could not be found")
		w.WriteHeader(http.StatusNotFound)
	})
}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = Handle404()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	fmt.Println("Server is up & running...")
	http.ListenAndServe(":3000", r)
}
