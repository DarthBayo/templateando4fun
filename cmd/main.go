package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			temp, err := template.ParseFiles("templates/index.html", "templates/blocks/home.html")
			if err != nil {
				log.Fatal(err)
			}

			err = temp.Execute(w, nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			temp, err := template.ParseFiles("templates/index.html", "templates/blocks/about.html")
			if err != nil {
				log.Fatal(err)
			}

			err = temp.Execute(w, nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
