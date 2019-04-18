package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// used to dump headers for debugging
func handler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	// disable cache
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// set hostname (used for demo)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "Error:", err)
	}

	// dump env
	for _, e := range os.Environ() {
		fmt.Fprintf(w, "%v\n", e)
	}
	fmt.Fprintf(w, "Served-By: %v\n", hostname)
	fmt.Fprintf(w, "Serving-Time: %s\n", time.Now().Sub(startTime))
	return

}

func main() {
	log.Print("env debug demo started")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
