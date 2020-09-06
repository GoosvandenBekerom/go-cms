package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port int `envconfig:"CMS_PORT" ,default:"8080" ,yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `envconfig:"CMS_DB_HOST" ,yaml:"host"`
		Port     int    `envconfig:"CMS_DB_PORT" ,yaml:"port"`
		Username string `envconfig:"CMS_DB_USERNAME" ,yaml:"username"`
		Password string `envconfig:"CMS_DB_PASSWORD" ,yaml:"password"`
		Dbname   string `envconfig:"CMS_DB_DBNAME" ,yaml:"dbname"`
	} `yaml:"database"`
}

var instance *Config = nil

func Get() *Config {
	if instance == nil {
		bootstrapConfigFile()
		bootstrapEnvironment()
	}
	return instance
}

func bootstrapConfigFile() {
	log.Println("Bootstrapping config.yaml file...")
	file, err := os.Open("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&instance)

	if err != nil {
		log.Fatal(err)
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
}

func bootstrapEnvironment() {
	log.Println("Bootstrapping environment variables...")
	if err := envconfig.Process("cms", instance); err != nil {
		log.Fatal(err)
	}
}
