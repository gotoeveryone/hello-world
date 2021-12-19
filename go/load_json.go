package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Id   int
	Name string
}

func main() {
	exe, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Errorf("Read executable error: %s", err))
		os.Exit(1)
	}

	jsonValue, err := ioutil.ReadFile(filepath.Join(exe, "../test.json"))
	if err != nil {
		fmt.Println(fmt.Errorf("Read file error: %s", err))
		os.Exit(1)
	}

	var config []Config
	if err := json.Unmarshal(jsonValue, &config); err != nil {
		fmt.Println(fmt.Errorf("Unmarshal error: %s", err))
		os.Exit(1)
	}
	for _, c := range config {
		fmt.Println(fmt.Sprintf("id: %d, name: %s", c.Id, c.Name))
	}
}
