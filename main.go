package main

import (
	"fmt"
	"net/http"
	"tikv-client/app"
)

func main() {
	// Static files (CSS, JS, images)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define routes
	http.HandleFunc("/get", app.RenderForm)
	http.HandleFunc("/getRecord", app.HandleFormSubmission)
	http.HandleFunc("/put", app.RenderPutForm)         // New route for the PUT form
	http.HandleFunc("/putRecord", app.HandlePutSubmission) // New route for handling PUT request

	// Start the server
	fmt.Println("Server started at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
