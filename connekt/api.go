package connekt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var apiEndpoint = os.Getenv("API_ENDPOINT")
var apiUrl = apiEndpoint + "/v2/send/email/"

var httpClient http.Client = http.Client{}

func SendEmail(request ConnektEmailRequest, appName string, apiKey string) (string, error) {
	messageId := ""
	b, err := json.Marshal(request)
	if err != nil {
		log.Printf("Error: %s", err)
		return messageId, err
	}

	req, err := http.NewRequest("POST", apiUrl+appName, bytes.NewBuffer(b))
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Send Request Error", err, "Request: ", string(b))
		return messageId, fmt.Errorf("Failed to Send Request to Connekt: %s", err)
	}
	defer resp.Body.Close()

	log.Println("Send Email Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("Send Email Response:", string(body))

	if resp.StatusCode/100 != 2 {
		var cerr ConnektErrorResponse
		err = json.Unmarshal(body, &cerr)
		if err != nil {
			return messageId, fmt.Errorf("Non2XX[%d] from Connekt", resp.StatusCode)
		}
		return messageId, fmt.Errorf("Non2XX[%d] from Connekt: %s", resp.StatusCode, cerr.Response.Message)
	} else {
		var jsonResp ConnektResponse
		err = json.Unmarshal(body, &jsonResp)
		if err != nil {
			log.Println("Response Deserialize Error", err)
			return messageId, nil // its okay, just json error
		}
		for k, _ := range jsonResp.Response.Success {
			messageId = k
			break
		}
		return messageId, nil
	}

}
