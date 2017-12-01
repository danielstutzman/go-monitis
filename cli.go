package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	ApiKey    string
	SecretKey string
}

func getAuthToken(apiKey, secretKey string) (string, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "http://www.monitis.com/api?action=authToken&apikey="+apiKey+"&secretkey="+secretKey, nil)
	if err != nil {
		return "", fmt.Errorf("Error from NewRequest: %s", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("Error from client.Do: %s", err)
	}

	defer response.Body.Close()
	responseJson := map[string]string{}
	err = json.NewDecoder(response.Body).Decode(&responseJson)
	if err != nil {
		return "", fmt.Errorf("Error from Decode: %s", err)
	}

	return responseJson["authToken"], nil
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

	log.Printf("Auth token: %s", authToken)
}
