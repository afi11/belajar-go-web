package main

import (
	"fmt"
	"net/http"

	// Tambahkan ini untuk latihan ke 3
	"html/template"
)

// func handlerIndex(w http.ResponseWriter, r *http.Request) {
// 	var message = "Welcome"
// 	w.Write([]byte(message))
// }

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}

// Tambah untuk step 4
type M map[string]interface{}

// Tambah untuk step 5
type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func main() {
	// http.HandleFunc("/", handlerIndex)
	// http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	// 2. Routing Static Assets
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	// 3. Render HTML File
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	var filepath = path.Join("views", "index.html")
	// 	var tmpl, err = template.ParseFiles(filepath)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	var data = map[string]interface{}{
	// 		"title": "Learning Golang Web",
	// 		"name":  "Batman",
	// 	}

	// 	err = tmpl.Execute(w, data)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
