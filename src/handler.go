package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handleCommandSetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	subset := r.PathValue("subset")
	id := r.PathValue("id")
	description := r.URL.Query().Get("description")

	commands := GetCommandsFromSubset(subset, id)
	if commands == nil {
		w.Write([]byte("No commands available at this ID"))
		return
	}

	fmt.Println(subset, id)
	fmt.Println(commands)

	var cmdStrings []string
	if description == "true" {
		cmdStrings = append(cmdStrings, "# "+commands.Description)
	}
	cmdStrings = append(cmdStrings, commands.Commands...)

	responseText := strings.Join(cmdStrings, "\n")

	// jsonBytes, err := json.Marshal(commands)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	w.Write([]byte(responseText + "\n"))
}

func handleAvailableSubsets(w http.ResponseWriter, r *http.Request) {
	DB, err := ReadCommandDB()
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	var subsets []string

	for _, s := range DB.CommandSubsets {
		subsets = append(subsets, s.Subset)
	}

	responseText := strings.Join(subsets, "\n")

	_, err = w.Write([]byte("# Available command datasets \n" + responseText + "\n"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func handleAvailableCommandsInSubset(w http.ResponseWriter, r *http.Request) {
	subset := r.PathValue("subset")

	DB, err := ReadCommandDB()
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "text/plain")

	var commandInSubsetNames []string
	for _, s := range DB.CommandSubsets {
		if s.Subset == subset {
			for _, cmd := range s.CommandsInSubset {
				commandInSubsetNames = append(commandInSubsetNames, cmd.ShortDescription)
			}
			responseText := strings.Join(commandInSubsetNames, "\n")
			w.Write([]byte("# Available commands in subset \n" + responseText + "\n"))
			return
		}
	}

	w.Write([]byte("No such subset found: " + subset + "\n"))
}
