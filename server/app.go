package main

import (
	"github.com/unrolled/render"
	"log"
	"net/http"
)

//
func main() {
	r := render.New()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "index", nil)
	})
	//// static files - js,css, img
	// in html /dist/js/jquery.min.js
	fs := http.FileServer(http.Dir("public"))

	// how we access these files
	// /static/dist/js/jquery.min.js
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening...")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
