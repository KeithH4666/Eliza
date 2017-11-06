package main

import (
	"net/http"
)

func main() {
	
	//Handles static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
    log.Println("Preparing guessing game , enter this in your web browser - Localhost:8080")
    http.ListenAndServe(":8080", nil)	
	
}