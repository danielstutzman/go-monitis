package main

import (
	"flag"
	"github.com/danielstutzman/go-monitis"
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

	authToken, err := monitis.GetAuthToken(config.ApiKey, config.SecretKey)
	if err != nil {
		log.Fatalf("Error from GetAuthToken: %s", err)
	}

	if false {
		alerts, err := monitis.GetRecentAlerts(config.ApiKey, authToken)
		if err != nil {
			log.Fatalf("Error from GetRecentAlerts: %s", err)
		}
		log.Printf("Recent alerts: %+v", alerts)
	}

}
