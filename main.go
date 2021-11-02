package main

import (
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	value, foundValue := os.LookupEnv(key)
	if foundValue {
		return value
	}
	return fallback
}

func main() {
	/*
	   Referenced to https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
	   for ways to load html files using go servers (Nov 1st, 2021)
	*/
	// Handle homepage and report pages
	http.Handle("/", http.FileServer(http.Dir("./reports")))
	// Start listening for HTTP requests.
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}
