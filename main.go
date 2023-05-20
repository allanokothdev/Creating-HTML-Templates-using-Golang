package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	pageTitle string
	Todos	[]Todo
}

func main() {
	tmp1 := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		data := TodoPageData{
			pageTitle: "My TODO list1",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmp1.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}