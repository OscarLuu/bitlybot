package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Short   string `json:"link"`
	Long    string `json:"long_url"`
	Created string `json:"created_at"`
}

func Bitly(link string) string {
	var response Response

	url := "https://api-ssl.bitly.com/v4/shorten"
	token := ""
	var bearer = "Bearer " + token

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

	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalf("Issue with unmarshal %s", err)
		}
		log.Printf("SUCCESS: Response code %d", resp.StatusCode)
		shortLink := response.Short
		return shortLink
	} else {
		log.Printf("ERROR: Response code %d", resp.StatusCode)
		return "error"
	}
}
