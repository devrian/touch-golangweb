package handler

import (
	"bwa-golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// HomeHandler is ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmp, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 11},
		{ID: 2, Name: "Xpander", Price: 230000000, Stock: 8},
		{ID: 3, Name: "Pajero Sport", Price: 240000000, Stock: 1},
	}

	err = tmp.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

// HelloHandler is ...
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, saya sedang belajar golang web"))
}

// MarioHandler is ...
func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from nintendo"))
}

// ProductHandler is ...
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"content": idNumb,
	}

	tmp, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

// PostGet is ...
func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}

// Form is ...
func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = temp.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}

// Process is ...
func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		temp, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
			return
		}

		err = temp.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}
