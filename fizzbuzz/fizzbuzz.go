package fizzbuzz

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type FizzbuzzCommand struct {
	Count *int `json:"count"`
}

type FizzbuzzResponse struct {
	Message string `json:"message"`
}

func getFizzbuzzString(count int) string {
	fizzbuzzStr := ""
	fizzStr := os.Getenv("FIZZ")
	buzzStr := os.Getenv("BUZZ")
	if count%3 == 0 {
		fizzbuzzStr += fizzStr
	}
	if count%5 == 0 {
		fizzbuzzStr += buzzStr
	}
	return fizzbuzzStr
}

func HandleFizzbuzz(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,POST")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.Header().Add("Allow", http.MethodOptions)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method " + r.Method + " not allowed"))
		return
	}
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("Content type " + contentType + " is not supported"))
		return
	}
	decoder := json.NewDecoder(r.Body)
	var command FizzbuzzCommand
	if err := decoder.Decode(&command); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error unmarshalling JSON: " + err.Error()))
		return
	}
	if command.Count == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON missing 'count' parameter"))
		return
	}
	replyMessage := FizzbuzzResponse{Message: getFizzbuzzString(*command.Count)}
	replyBytes, err := json.Marshal(replyMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshalling JSON: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(replyBytes))
}
