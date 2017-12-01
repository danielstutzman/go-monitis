package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
