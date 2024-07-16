package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Use:", os.Args[0], "<domain> <api-key>")
		os.Exit(0)
	}

	domain := os.Args[1]
	url := "https://www.virustotal.com/api/v3/domains/" + domain + "/subdomains?limit=1966"
	headers := map[string]string{
		"accept":   "application/json",
		"x-apikey": os.Args[2],
	}

	body, err := httpGet(url, headers)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(body), &jsonMap)

	dataArray := jsonMap["data"].([]interface{})

	for i := 0; i < len(dataArray); i++ {
		subdomains := dataArray[i].(map[string]interface{})
		id := subdomains["id"].(string)
		fmt.Println(id)
	}

}

func httpGet(url string, headers map[string]string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return "", fmt.Errorf("error: Status Code %d, verify API-KEY", res.StatusCode)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: Status Code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
