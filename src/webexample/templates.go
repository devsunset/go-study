package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)

	// {{/* a comment */}}	Defines a comment
	// {{.}}	Renders the root element
	// {{.Title}}	Renders the “Title”-field in a nested element
	// {{if .Done}} {{else}} {{end}}	Defines an if-Statement
	// {{range .Todos}} {{.}} {{end}}	Loops over all “Todos” and renders each using {{.}}
	// {{block "content" .}} {{end}}	Defines a block with the name “content”
}
