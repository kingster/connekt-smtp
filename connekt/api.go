package connekt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var apiEndpoint = os.Getenv("API_ENDPOINT")
var apiUrl = apiEndpoint + "/v2/send/email/"

var httpClient http.Client = http.Client{
	Timeout: 10 * time.Second,
}

type SendMailResult struct {
	Status       int
	MessageId    string
	ErrorMessage string
}

func SendEmail(request ConnektEmailRequest, appName string, apiKey string) (*SendMailResult, error) {
	result := &SendMailResult{}
	b, err := json.Marshal(request)
	if err != nil {
		log.Printf("Error: %s", err)
		return result, err
	}

	req, err := http.NewRequest("POST", apiUrl+appName, bytes.NewBuffer(b))
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Send Request Error", err, "Request: ", string(b))
		return result, fmt.Errorf("Failed to Send Request to Connekt: %s", err)
	}
	defer resp.Body.Close()

	log.Println("Send Email Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("Send Email Response:", string(body))

	if resp.StatusCode/100 != 2 {
		var cerr ConnektErrorResponse
		err = json.Unmarshal(body, &cerr)
		if err != nil {
			return result, fmt.Errorf("Non2XX[%d] from Connekt", resp.StatusCode)
		} else {
			result.ErrorMessage = cerr.Response.Message
		}
		return result, fmt.Errorf("Non2XX[%d] from Connekt: %s", resp.StatusCode, cerr.Response.Message)
	} else {
		var jsonResp ConnektResponse
		err = json.Unmarshal(body, &jsonResp)
		if err != nil {
			log.Println("Response Deserialize Error", err)
			return result, nil // its okay, just json error
		}
		for k := range jsonResp.Response.Success {
			result.MessageId = k
			break
		}
		result.Status = jsonResp.Status
		return result, nil
	}

}
