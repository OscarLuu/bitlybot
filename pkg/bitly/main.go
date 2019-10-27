package bitly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const SHORTENURL = "https://api-ssl.bitly.com/v4/shorten"

var (
	std = New()
)

type BitlyAPI struct {
	token string
	cli   *http.Client
}

func New() *BitlyAPI {
	return &BitlyAPI{
		cli: http.DefaultClient,
	}
}

func SetToken(token string) {
	std.token = token
}

type Response struct {
	Short   string `json:"link"`
	Long    string `json:"long_url"`
	Created string `json:"created_at"`
}

func Shorten(link string) (string, error) {
	return std.Shorten(link)
}

func (bapi BitlyAPI) Shorten(link string) (string, error) {

	data := map[string]string{"long_url": link}
	postBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := createRequest(postBody)
	if err != nil {
		return "", err
	}

	resp, err := bapi.cli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return "", fmt.Errorf("status code %v", resp.StatusCode)
	}

	responseBody := Response{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.Short, nil
}

func createRequest(b []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", SHORTENURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", std.token))
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

/*
func Bitly(link string) string {
	var response Response

	token := ""
	var bearer = "Bearer " + token

	jsonData := map[string]string{"long_url": link}
	jsonValue, _ := json.Marshal(jsonData)

	req, err := http.NewRequest("POST", SHORTENURL, bytes.NewBuffer(jsonValue))
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
*/
