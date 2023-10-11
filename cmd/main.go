package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			temp, err := template.ParseFiles("templates/index.html")
			if err != nil {
				log.Fatal(err)
			}

			err = temp.Execute(w, nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/secret-content", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("Now we have some content"))
		}
	})

	http.HandleFunc("/other-content", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("Welcome to where time stands still"))
		}
	})

	http.HandleFunc("/1s-delay", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("Here we have some content with 1s delay"))
		}
	})

	http.HandleFunc("/7s-delay", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("I made this at night time, after a very stressfull day. Again, say bye to creativity. I'm just trying to have some fun"))
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
