package config

import (
"gopkg.in/yaml.v2"
"log"
)

//var appConfigs configs.AppConfig

func init() {
	c := app.AppConfig{}
	err := yaml.Unmarshal([]byte(""), &c)
	if err != nil {
		log.Fatalf("Configurations load failed: %v", err)
	}

	log.Printf("Configurations file loaded %s", c)
}

func Get() {

}

func loadDefaults() {

}
