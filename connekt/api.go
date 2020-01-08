package connekt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var apiEndpoint = os.Getenv("API_ENDPOINT")

func SendEmail(request ConnektEmailRequest, apiKey string) error {

	b, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(b))
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Send Request Error", err)
	}
	defer resp.Body.Close()

	fmt.Println("Send Email Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Send Email Response:", string(body))

	//var jsonResp ConnektResponse
	//err = json.Unmarshal(body, &jsonResp)
	//if err != nil {
	//	fmt.Println("Response Deserlize Error", err)
	//}

	if resp.StatusCode /100 != 2 {
		var cerr ConnektErrorResponse
		err = json.Unmarshal(body, &cerr)
		if err != nil {
			return fmt.Errorf("Non2XX[%d] from Connekt", resp.StatusCode)
		}
		return fmt.Errorf("Non2XX[%d] from Connekt: %s", resp.StatusCode, cerr.Response.Message)
	}

	return nil
}
