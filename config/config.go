package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Cities []City
}

type City struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = _new()
	}
	return config
}

func _new() *Config {
	// get file content from env
	content, err := os.ReadFile(os.Getenv("COVID_INCIDENCE_CONFIG"))
	check(err)
	//fmt.Print(string(content))

	// parse content file
	var data map[string]string
	err = json.Unmarshal(content, &data)
	check(err)

	var c Config
	for name, url := range data {
		c.Cities = append(c.Cities, City{Name: name, URL: url})
	}

	//fmt.Println(c)

	return &c
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
