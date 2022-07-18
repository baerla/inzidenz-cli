package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Cities []City
}

type City struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

func (c *Config) UnmarshalJSON(b []byte) error {
	var data map[string]string
	err := json.Unmarshal(b, &data)
	check(err)

	for name, url := range data {
		c.Cities = append(c.Cities, City{Name: name, URL: url})
	}
	return nil
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = importFromConfigFile()
	}
	return config
}

func importFromConfigFile() *Config {
	// get file content from env
	content, err := os.ReadFile(os.Getenv("COVID_INCIDENCE_CONFIG"))
	check(err)

	// parse content file
	var c Config
	err = json.Unmarshal(content, &c)
	check(err)

	//fmt.Println(c)

	return &c
}

func (c *Config) SaveToConfigFile() {
	// export config to file
	data := make(map[string]string)
	for _, city := range c.Cities {
		data[city.Name] = city.URL
	}
	content, err := json.Marshal(data)
	check(err)
	err = ioutil.WriteFile(os.Getenv("COVID_INCIDENCE_CONFIG"), content, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
