package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var PORT, WEBHOOK_URL string

type InitiateResponse struct {
	ResumeURL string `json:"resumeURL"`
	Mode      string `json:"mode"`
	ID        string `json:"id"`
	FormURL   string `json:"formUrl"`
}

// CORS middleware
func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

// handleInitiate calls the n8n "initiate" webhook
func handleInitiate(w http.ResponseWriter, r *http.Request) {

	req, err := http.NewRequest("POST", WEBHOOK_URL, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20) // max 10MB

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to read file: "+err.Error(), 400)
		return
	}
	defer file.Close()

	resumeURL := r.FormValue("resumeURL")
	if resumeURL == "" {
		http.Error(w, "Missing resumeURL", 400)
		return
	}

	action := r.FormValue("action")
	if action == "" {
		action = "something" // default action if none provided
	}

	// ---- Build multipart/form-data request ----
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the action field
	writer.WriteField("action", action)

	// Add the file field
	part, err := writer.CreateFormFile("image", handler.Filename)
	if err != nil {
		http.Error(w, "Failed to create form file: "+err.Error(), 500)
		return
	}
	io.Copy(part, file)
	writer.Close()

	// Send multipart request to n8n
	req, err := http.NewRequest("POST", resumeURL, body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to send to n8n: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	WEBHOOK_URL = os.Getenv("N8N_WEBHOOK_URL")
	fmt.Println(WEBHOOK_URL)
}

func main() {

	http.HandleFunc("/initiate", withCORS(handleInitiate))
	http.HandleFunc("/upload", withCORS(handleUpload))

	fmt.Println("Server running on http://localhost:" + PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		fmt.Println("Server error:", err)
		os.Exit(1)
	}
}
