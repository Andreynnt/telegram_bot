package config

import (
	"fmt"
	"os"
	"encoding/json"
)

type Settings struct {
	Token string

}

func Read(path string) (Settings, error) {
	var settings Settings
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		fmt.Println("Can't open file\n")
		return settings, err
	}

	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&settings); err != nil {
		fmt.Println("Can't jsonParser.Decode\n")
		return settings, err
	}

	return settings, err
}