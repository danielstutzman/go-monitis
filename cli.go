package main

import (
	"flag"
	"log"
)

type Config struct {
	ApiKey    string
	SecretKey string
}

func main() {
	config := Config{}
	flag.StringVar(&config.ApiKey, "apikey", "", "API key for Monitis")
	flag.StringVar(&config.SecretKey, "secretkey", "", "Secret key for Monitis")
	flag.Parse()

	if config.ApiKey == "" {
		log.Fatalf("Missing -apikey")
	}
	if config.SecretKey == "" {
		log.Fatalf("Missing -secretkey")
	}

	authToken, err := getAuthToken(config.ApiKey, config.SecretKey)
	if err != nil {
		log.Fatalf("Error from getAuthToken: %s", err)
	}

	alerts, err := getRecentAlerts(config.ApiKey, authToken)
	if err != nil {
		log.Fatalf("Error from getRecentAlerts: %s", err)
	}

	log.Printf("Recent alerts: %+v", alerts)
}