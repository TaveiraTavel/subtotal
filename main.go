package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/TaveiraTavel/subtotal/static"
)

func main() {
	domains := []string{}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		domains = append(domains, stdin.Text())
	}

	static.ShowBanner()

	for _, domain := range domains {
		virusTotalSeach(domain)
	}
}

func virusTotalSeach(domain string) {
	url := "https://www.virustotal.com/api/v3/domains/" + fmt.Sprint(domain) + "/subdomains?limit=1966"
	headers := map[string]string{
		"accept":   "application/json",
		"x-apikey": os.Getenv("VT_API_KEY"),
	}

	body, err := httpGet(url, headers)
	if err != nil {
		fmt.Println(err)
		return
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(body), &jsonMap)

	dataArray := jsonMap["data"].([]interface{})

	for i := 0; i < len(dataArray); i++ {
		subdomains := dataArray[i].(map[string]interface{})
		id := subdomains["id"].(string)
		fmt.Println(id)
	}

	return
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
		return "", fmt.Errorf("error: %d Unauthorized\nVerify VT_API_KEY Enviroment Variable", res.StatusCode)
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

