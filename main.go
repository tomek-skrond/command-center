package main

import (
	"log"
	"net/http"
)

//	func handleCommandByID(w http.ResponseWriter, r *http.Request) {
//		return
//	}

var DB = CommandDB{}

func main() {
	r := http.NewServeMux()

	// r.HandleFunc("GET /command/{id}", handleCommandByID)
	r.HandleFunc("GET /command/{subset}/{id}", handleCommandSetByID)
	r.HandleFunc("GET /available/subset", handleAvailableSubsets)
	r.HandleFunc("GET /available/subset/{subset}", handleAvailableCommandsInSubset)

	log.Fatalln(http.ListenAndServe(":2137", r))
}
