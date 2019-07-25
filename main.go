package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

/*
Hello defines the two things that get displayed in the front.
*/
type Hello struct {
	Name string
	Time string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	hello := Hello{"Nathan", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("template/hello-world.html"))

	if name := r.FormValue("name"); name != "" {
		hello.Name = name
	}

	if err := templates.ExecuteTemplate(w, "hello-world.html", hello); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	fmt.Println("Listening on port 8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
