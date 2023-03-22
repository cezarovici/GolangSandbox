package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type VonageResponse struct {
	Messages []struct {
		Status    string `json:"status"`
		ErrorText string `json:"error-text"`
	} `json:"messages"`
}

func NewRequest(apiPath string, body url.Values) (*http.Request, error) {
	return http.NewRequest("POST", apiPath, strings.NewReader(body.Encode()))
}

const (
	apiKey    = "<apiKey>"
	apiSecret = "<apiSecret>"
	apiPath   = "https://rest.nexmo.com/sms/json"

	from = "User-name"
	to   = "<Enter-receive-number>"

	message = "This is a message send from Voyage API"
)

var body = url.Values{
	"from":       {from},
	"to":         {to},
	"text":       {message},
	"api_key":    {apiKey},
	"api_secret": {apiSecret},
}

func main() {
	request, errNewRequest := NewRequest(apiPath, body)
	if errNewRequest != nil {
		fmt.Print(errNewRequest)
		os.Exit(1)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, errDoRequest := client.Do(request)
	if errDoRequest != nil {
		fmt.Print(errDoRequest)
		os.Exit(2)
	}

	defer res.Body.Close()

	resp := &VonageResponse{}

	derr := json.NewDecoder(res.Body).Decode(resp)
	if derr != nil {
		fmt.Println(derr)
		os.Exit(3)
	}

	if len(resp.Messages) <= 0 {
		fmt.Println("Vonage error: Internal Error")
		os.Exit(4)
	}

	if resp.Messages[0].Status != "0" {
		fmt.Printf("Vonage error: %v (status: %v)", resp.Messages[0].ErrorText, resp.Messages[0].Status)
		os.Exit(5)
	}

	fmt.Println("SMS sent successfully.")
}
