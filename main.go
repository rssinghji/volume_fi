package main

import (
	"log"
	"net/http"
	"runtime"

	api "volumefi.com/volume_fi/api"
)

func main() {
	log.Println("Entrypoint/Start for the exam question")
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	// Required API
	http.HandleFunc("/", api.HandleDefault)
	http.HandleFunc("/calculate", api.HandleCalculateFlightTracker)

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
