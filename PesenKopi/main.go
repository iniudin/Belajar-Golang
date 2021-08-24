package main

import (
	"html/template"
	"net/http"
	"path"
)

func index(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "index.html")
	template, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title": "PesenKopi - yok pesen kopinya bang!",
		"name":  "Pelanggan",
	}

	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "index.html")
	template, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title": "PesenKopi - yok pesen kopinya bang!",
		"name":  "Pelanggan",
	}

	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	template, err := template.ParseGlob("views/**")
	if err != nil {
		panic(err.Error())
		return
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
