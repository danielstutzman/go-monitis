package monitis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type GetExternalResultsOutput struct {
	LocationName string          `json:"locationName"` // e.g. "USA-WST"
	DataTuples   [][]interface{} `json:"data"`
	Trend        PointsTrend     `json:"trend"`
	Points       []Point
}

type Point struct {
	Timestamp time.Time
	Duration  float64
	WasOkay   bool
}

type PointsTrend struct {
	Min        float64 `json:"min"`
	OkCount    int     `json:"okcount"`
	Max        float64 `json:"max"`
	OkSum      float64 `json:"oksum"`
	NotOkCount float64 `json:"nokcount"`
}

func (auth *Auth) GetExternalResults(testId string) ([]GetExternalResultsOutput,
	error) {

	client := &http.Client{}

	request, err := http.NewRequest("GET", "http://www.monitis.com/api"+
		"?action=testresult"+
		"&testId="+testId+
		"&apikey="+auth.ApiKey+
		"&authToken="+auth.AuthToken, nil)
	if err != nil {
		return []GetExternalResultsOutput{},
			fmt.Errorf("Error from NewRequest: %s", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return []GetExternalResultsOutput{},
			fmt.Errorf("Error from client.Do: %s", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if body[0] == '{' {
		output := map[string]interface{}{}
		err = json.Unmarshal(body, &output)
		if err != nil {
			return []GetExternalResultsOutput{},
				fmt.Errorf("Error from Unmarshal as object: %s", err)
		}
		return []GetExternalResultsOutput{},
			fmt.Errorf("API responded with error: %s", output["error"])
	} else if body[0] == '[' {
		output := []GetExternalResultsOutput{}
		err = json.Unmarshal(body, &output)
		if err != nil {
			return []GetExternalResultsOutput{},
				fmt.Errorf("Error from Unmarshal as array: %s", err)
		}

		newOutput := []GetExternalResultsOutput{}
		for _, result := range output {
			for _, tuple := range result.DataTuples {
				timestamp, err := time.Parse("2006-01-02 15:04", tuple[0].(string))
				if err != nil {
					return []GetExternalResultsOutput{},
						fmt.Errorf("Can't decode timestamp '%s': %s", tuple[0], err)
				}

				point := Point{
					Timestamp: timestamp,
					Duration:  tuple[1].(float64),
					WasOkay:   tuple[2].(string) == "OK",
				}
				result.Points = append(result.Points, point)
			}
			result.DataTuples = [][]interface{}{}
			newOutput = append(newOutput, result)
		}
		return newOutput, nil
	} else {
		return []GetExternalResultsOutput{},
			fmt.Errorf("API responded with unexpected first character: %s", body)
	}
}
