package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Bitly takes the auth token and a url string and returns a shortened
// link
func Bitly(link string) {
	url := "https://api-ssl.bitly.com/v4/shorten"
	token := "<redact>"
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + token

	// Create a new request using http
	// Json marshall data
	jsonData := map[string]string{"long_url": link}
	jsonValue, _ := json.Marshal(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalf("Error with setting up request")
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error on response")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))
}
