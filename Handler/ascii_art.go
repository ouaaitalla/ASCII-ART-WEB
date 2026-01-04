package handel

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"text/template"

)

func Ascii_art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad reauest", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	if len(text) > 1024 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if text == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	text = strings.ReplaceAll(text, "\r\n", "\n")
	if !IsValidASCII(text) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	banner := r.FormValue("banner")
	if banner == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	bannerx, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		http.Error(w, "internal sever error", http.StatusInternalServerError)
		return
	}
	result := GenerateAscii(text, bannerx)
	
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	var buf bytes.Buffer
	err1 := tmpl.Execute(&buf, result)
	if err1 != nil {
		http.Error(w, "internal sever error", 500)
		return
	}
	w.WriteHeader(200)
	buf.WriteTo(w)
}
