package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"

	"relay/templates"
)

func main() {
	token, chat := os.Getenv("TELEGRAM_BOT_TOKEN"), os.Getenv("TELEGRAM_CHAT_ID")
	if token == "" || chat == "" {
		log.Fatal("consider filling out the env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	send := func(w http.ResponseWriter, msg, mode string) {
		v := url.Values{"chat_id": {chat}, "text": {msg}}
		if mode != "" {
			v.Set("parse_mode", mode)
		}
		r, err := http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token), v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer r.Body.Close()
		w.WriteHeader(r.StatusCode)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		var b struct{ Message string `json:"message"` }
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil || b.Message == "" {
			http.Error(w, `expected JSON {"message":"..."}`, http.StatusBadRequest)
			return
		}
		send(w, b.Message, "")
	})

	for path, tmpl := range templates.All {
		t := template.Must(template.New(path).Parse(tmpl))
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "POST only", http.StatusMethodNotAllowed)
				return
			}
			var d map[string]any
			if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
				http.Error(w, "bad JSON", http.StatusBadRequest)
				return
			}
			var buf bytes.Buffer
			if err := t.Execute(&buf, d); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			send(w, buf.String(), "Markdown")
		})
	}

	log.Printf("relay listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
