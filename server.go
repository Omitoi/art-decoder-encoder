package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func runServer() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmplParsed := template.Must(template.ParseFiles("template.html"))
	data := struct{ Result string }{Result: ""}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "<span style=\"color:rgb(255, 0, 0)\">Error parsing form</span>", http.StatusBadRequest)
			return
		}
		input := strings.Split(r.FormValue("art_input"), "\n")
		action := r.FormValue("action")
		var correct bool
		var result []string
		if len(input) < 1 || (len(input) == 1 && input[0] == "") {
			w.WriteHeader(http.StatusBadRequest)
			data.Result = "<span style=\"color:rgb(255, 0, 0)\">Please submit any input to the textarea above</span>"
			tmplParsed.Execute(w, data)
			return
		}
		println(input[0])
		if strings.ToLower(action) == "encode" {
			result, correct = encoder(input)
		} else if strings.ToLower(action) == "decode" {
			result, correct = decoder(input, "")
		}
		if !correct {
			w.WriteHeader(http.StatusBadRequest)
			data.Result = "<span style=\"color:rgb(255, 0, 0)\">Something went wrong</span>"
		} else {
			w.WriteHeader(http.StatusAccepted)
			data.Result = strings.Join(result, "\n")
		}
	}

	tmplParsed.Execute(w, data)
}
