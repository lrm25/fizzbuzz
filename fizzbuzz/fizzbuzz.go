package fizzbuzz

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const ENV_FIZZ = "FIZZ"
const ENV_BUZZ = "BUZZ"

type FizzbuzzCommand struct {
	Count *int `json:"count"`
}

type FizzbuzzResponse struct {
	Message string `json:"message"`
}

type FizzbuzzServer struct {
	port      string
	clientUrl string
}

func NewFizzbuzzServer(port, clientUrl string) *FizzbuzzServer {
	return &FizzbuzzServer{
		port:      port,
		clientUrl: clientUrl,
	}
}

func (f *FizzbuzzServer) Start() {
	http.HandleFunc("/fizzbuzz", f.handleFizzbuzz)
	if err := http.ListenAndServe(":"+f.port, nil); err != nil {
		panic("Error starting server on port " + f.port + ": " + err.Error())
	}
}

func (f *FizzbuzzServer) getFizzbuzzString(count int) string {
	fizzbuzzStr := ""
	fizzStr := os.Getenv(ENV_FIZZ)
	buzzStr := os.Getenv(ENV_BUZZ)
	if count%3 == 0 {
		fizzbuzzStr += fizzStr
	}
	if count%5 == 0 {
		fizzbuzzStr += buzzStr
	}
	return fizzbuzzStr
}

func (f *FizzbuzzServer) handleError(w http.ResponseWriter, errCode int, errString string) {
	fmt.Println(errString)
	w.WriteHeader(errCode)
	if _, err := w.Write([]byte(errString)); err != nil {
		fmt.Println("Error writing error data to response: " + err.Error())
	}
}

func (f *FizzbuzzServer) handleFizzbuzz(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", f.clientUrl)
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,POST")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.Header().Add("Allow", http.MethodOptions)
		f.handleError(w, http.StatusMethodNotAllowed, "Method "+r.Method+" not allowed")
		return
	}
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		f.handleError(w, http.StatusUnsupportedMediaType, "Content type "+contentType+" is not supported")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var command FizzbuzzCommand
	if err := decoder.Decode(&command); err != nil {
		f.handleError(w, http.StatusBadRequest, "Error unmarshalling JSON: "+err.Error())
		return
	}
	if command.Count == nil {
		f.handleError(w, http.StatusBadRequest, "JSON missing 'count' parameter")
		return
	}
	replyMessage := FizzbuzzResponse{Message: f.getFizzbuzzString(*command.Count)}
	replyBytes, err := json.Marshal(replyMessage)
	if err != nil {
		f.handleError(w, http.StatusInternalServerError, "Error marshalling JSON: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(replyBytes)); err != nil {
		fmt.Println("Error writing reply to response: " + err.Error())
	}
}
