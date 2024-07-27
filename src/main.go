package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//	func handleCommandByID(w http.ResponseWriter, r *http.Request) {
//		return
//	}

var DB = CommandDB{}

func apiVersion1(r *http.ServeMux) {
	r.HandleFunc("GET /command/{subset}/{id}", handleCommandSetByID)
	r.HandleFunc("GET /available/subset", handleAvailableSubsets)
	r.HandleFunc("GET /available/subset/{subset}", handleAvailableCommandsInSubset)
}

func apiVersion2(r *http.ServeMux) {
	r.HandleFunc("GET /v2/command/subset/{subset}/{id}", handleCommandSetByID)
	r.HandleFunc("GET /v2/command/subset/{subset}", handleAvailableCommandsInSubset)
	r.HandleFunc("GET /v2/command/subset", handleAvailableSubsets)
}

func initDatabase() (*PostgresStorage, error) {
	testenv := os.Getenv("ENV")

	if testenv == "test" {
		s, err := NewPostgresStorage(WithTestEnv())
		if err != nil {
			return nil, err
		}
		return s, nil
	}

	if testenv != "prod" {
		s, err := NewPostgresStorage(WithProdEnv())
		if err != nil {
			return nil, err
		}
		return s, nil
	}

	return nil, fmt.Errorf("no such environment")
}
func main() {

	s, err := initDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)

	r := http.NewServeMux()

	apiVersion1(r)
	apiVersion2(r)

	log.Fatalln(http.ListenAndServe(":2137", r))
}
