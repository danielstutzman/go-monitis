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

	if false {
		m := monitis.AddExternalMonitorOptions{
			Name:             monitis.String("myname"),
			DetailedTestType: monitis.Int(1),
			Tag:              monitis.String("Default"),
			LocationIds:      monitis.String("1,9"),
			Url:              monitis.String("www.example.com"),
			Type:             monitis.String("http"),
			Interval:         monitis.Int(15),
		}
		monitor, err := monitis.AddExternalMonitor(&m, config.ApiKey, authToken)
		if err != nil {
			log.Fatalf("Error from AddExternalMonitor: %s", err)
		}
		log.Printf("New monitor: %+v", monitor)
	}

	if false {
		monitors, err := monitis.GetExternalMonitors(config.ApiKey, authToken)
		if err != nil {
			log.Fatalf("Error from GetExternalMonitors: %s", err)
		}
		log.Printf("External monitors: %+v", monitors)
	}

	if false {
		testId := "773757"
		results, err := monitis.GetExternalResults(testId, config.ApiKey, authToken)
		if err != nil {
			log.Fatalf("Error from GetExternalResults: %s", err)
		}
		log.Printf("External results: %+v", results)
	}
}
