package main

import (
	"fmt"
	"net/http"
	"time"
	"html/template"
)

type Hello struct {
	Name string
	Time string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	hello := Hello{"Nathan", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("template/hello-world.html"))

	if name := r.FormValue("name"); name != "" {
		hello.Name = name;
	}

	if err := templates.ExecuteTemplate(w, "hello-world.html", hello); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", handleRequest)
	fmt.Println("Listening on port 8080");
    fmt.Println(http.ListenAndServe(":8080", nil));
}