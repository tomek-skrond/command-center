package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	uuid "github.com/google/uuid"
)

type CommandSet struct {
	CommandSetID     uuid.UUID `gorm:"primaryKey" json:"id"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	Commands         []string  `json:"commands"`
}

type CommandSubset struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	SubsetName string    `gorm:"foreignKey:CommandSetID" json:"subset_name"`
}

type CommandSubsetDB struct {
	Subset           string
	CommandsInSubset []*CommandSet
}

type CommandDB struct {
	CommandSubsets []CommandSubsetDB
}

func ReadCommandDB() (*CommandDB, error) {

	commandDB := new(CommandDB)

	subsetArr, err := os.ReadDir("commands/")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, cmdsubset := range subsetArr {
		path := fmt.Sprintf("commands/%s", cmdsubset.Name())
		jsonFile, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer jsonFile.Close()

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		var cmds []*CommandSet
		// Unmarshal the JSON data into the CommandSet structure
		if err := json.Unmarshal(jsonBytes, &cmds); err != nil {
			log.Println(err)
			return nil, err
		}

		//assign all commands in subset to SubsetDB
		newCmdSubset := CommandSubsetDB{
			Subset:           strings.TrimSuffix(cmdsubset.Name(), ".json"),
			CommandsInSubset: cmds,
		}
		commandDB.CommandSubsets = append(commandDB.CommandSubsets, newCmdSubset)
	}

	return commandDB, nil
}
