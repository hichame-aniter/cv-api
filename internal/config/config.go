package config

import (
	"log"
	"os"
)

var DataPath string
var Port string
var Environment string // "development", "staging", "production"

func LoadConfig() {
	path, found := os.LookupEnv("DATA_PATH")
	if !found {
		path = "data/cv.json"
	}
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8080" // default port
	}
	env, found := os.LookupEnv("ENV")
	if !found {
		env = "dev" // default environment
	}

	log.Println("Server running on port:", port)
	log.Println("Using data file:", path)
	log.Printf("Running a '%s' environment.", env)
	DataPath = path
	Port = port
	Environment = env
}
