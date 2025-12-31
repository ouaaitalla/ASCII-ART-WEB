package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/ascii-art", ascii_art)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "bad request", 400)
	}
	if r.URL.Path != "/" {
		http.Error(w, "ERROR: THAT NOT FOUND", 404)
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return
	}
	tmpl.Execute(w, nil)
}

func ascii_art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad reauest", 400)
	}
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "bad reauest", 400)
		return
	}
	text := r.FormValue("text")
	text = strings.ReplaceAll(text, "\r\n", "\n")
	if !isValidASCII(text) {
		http.Error(w, "internal sever: 500", 500)
		return
	}
	banner := r.FormValue("banner")
	bannerx, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		
		http.Error(w, "internal sever: 500", 500)
		return
	}
	result := generateAscii(text, bannerx)
	res := strings.Join(result, "")
	data := map[string]string{
		"Result": res,
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "template not found", 404)
	}
	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println(err)
		http.Error(w, "INTERNAL SERVER ERROR", 500)
	}
}
