package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// disable cache
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// say hello
	w.Write([]byte("Hello there!"))
	return

}

// used to dump headers for debugging
func killswitchHandler(w http.ResponseWriter, r *http.Request) {

	// disable cache
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// return
	w.WriteHeader(200)
	w.Write([]byte("OK"))

	// https://shapeshed.com/unix-exit-codes/
	os.Exit(0)

}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/killswitch", killswitchHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
