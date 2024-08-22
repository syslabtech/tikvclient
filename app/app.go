package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var apiURL = func() string {
	if url := os.Getenv("TIKV_URL"); url != "" {
		return url
	}
	return "http://192.168.26.64:30901/"
}()

// RenderForm renders the form for the POST request
func RenderForm(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "getRecord.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// HandleFormSubmission handles the form submission for POST request
func HandleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		project := r.FormValue("project")
		key := r.FormValue("key")

		data := RequestData{
			Project: project,
			Key:     key,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
			return
		}

		apiURL := apiURL + "kvapi/getSingleRecord"
		resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Error making request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response", http.StatusInternalServerError)
			return
		}
		log.Printf("Response from server: %s", string(body))
		fmt.Fprintf(w, "Response from server: %s", string(body))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// RenderPutForm renders the form for the PUT request
func RenderPutForm(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "putRecord.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// HandlePutSubmission handles the form submission for PUT request
func HandlePutSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		project := r.FormValue("project")
		key := r.FormValue("key")
		value := r.FormValue("value")

		data := PutRequestData{
			Project: project,
			Key:     key,
			Value:   value,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
			return
		}

		apiURL := apiURL + "kvapi/putSingleRecord"
		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Error creating request", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error making request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response", http.StatusInternalServerError)
			return
		}
		log.Printf("Response from server: %s", string(body))
		fmt.Fprintf(w, "Response from server: %s", string(body))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
