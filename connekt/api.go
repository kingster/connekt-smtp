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
var apiKey = os.Getenv("API_KEY")

func SendEmail(request ConnektEmailRequest) {

	b, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
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
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	//var jsonResp ConnektResponse
	//err = json.Unmarshal(body, &jsonResp)
	//if err != nil {
	//	fmt.Println("Response Deserlize Error", err)
	//}

	fmt.Println("Send Email Response:", string(body))


}
