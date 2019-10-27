package bitly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const SHORTENURL = "https://api-ssl.bitly.com/v4/shorten"

var (
	// std is the default instance of BitlyAPI.
	std = New()
)

// BitlyAPI
type BitlyAPI struct {
	token string
	cli   *http.Client
}

// New creates a new BitlyAPI with a default HTTP client.
func New() *BitlyAPI {
	return &BitlyAPI{
		cli: http.DefaultClient,
	}
}

// SetToken sets the token for the BitlyAPI.
func SetToken(token string) {
	std.token = token
}

type Response struct {
	Short   string `json:"link"`
	Long    string `json:"long_url"`
	Created string `json:"created_at"`
}

// Shorten calls bapi.Shorten
func Shorten(link string) (string, error) {
	return std.Shorten(link)
}

// Shorten shortens links by calling the Bitly API.
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

// createRequest creates a POST request.
func createRequest(b []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", SHORTENURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", std.token))
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}
