package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// Payload represents the message payload. It can be JSON, form data or plain text. It's intended to be used with Notify.
type Payload struct {
	Subject string `json:"subject"`
	Body    string `json:"message"`
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Request body must not be empty
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	// Decode the request body based on the content type
	var payload Payload
	contentType := r.Header.Get("Content-Type")

	switch contentType {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "application/x-www-form-urlencoded":
		payload.Subject = r.FormValue("subject")
		payload.Body = r.FormValue("message")
	case "text/plain":
		payloadRaw, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		payload.Body = string(payloadRaw)
	default:
		http.Error(w, "unsupported content type: "+contentType, http.StatusBadRequest)
	}

	// Basic validation
	if payload.Subject == "" && payload.Body == "" {
		http.Error(w, "empty message", http.StatusBadRequest)
		return
	}

	// Set together the message
	var message string
	if payload.Subject != "" {
		message = payload.Subject + " - "
	}
	message += payload.Body

	// Log the message
	log.Println(message)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handleMessage)

	address := ":8080"
	if os.Getenv("NOTIFY_TEST_ADDRESS") != "" {
		address = os.Getenv("NOTIFY_TEST_ADDR")
	}

	log.Printf("Listening on %s", address)
	log.Fatalln(http.ListenAndServe(address, nil))
}
