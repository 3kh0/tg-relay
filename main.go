package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	token, chatID := os.Getenv("TELEGRAM_BOT_TOKEN"), os.Getenv("TELEGRAM_CHAT_ID")
	if token == "" || chatID == "" {
		log.Fatal("consider filling out the env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		var body struct{ Message string `json:"message"` }
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Message == "" {
			http.Error(w, "expected JSON {\"message\":\"...\"}", http.StatusBadRequest)
			return
		}
		resp, err := http.PostForm(
			fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token),
			url.Values{"chat_id": {chatID}, "text": {body.Message}},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		w.WriteHeader(resp.StatusCode)
	})

	log.Printf("relay listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
