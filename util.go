package main

import "log"

func GetCommandsFromSubset(subset string, id string) *CommandSet {
	DB, err := NewCommandDB()
	if err != nil {
		log.Println(err)
		return nil
	}
	for _, x := range DB.CommandSubsets {
		if x.Subset == subset {
			for _, y := range x.CommandsInSubset {
				if y.ShortDescription == id {
					return y
				}
			}
		}
	}
	return nil
}
