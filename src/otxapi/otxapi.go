package otxapi

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

const (
	get = "GET"
	post = "POST"
	baseURL = "https://otx.alienvault.com"

	subscriptionsURLPath = "/api/v1/pulses/subscribed"
	APIKey = "db91e98e6dcac6303bd1522d3542f24fcb4be176ea262ecd892d39e0d82a218b"

)

type SubscribedPulses struct {
    data map[string]interface{} `json:"results"`
}

func (c *SubscribedPulses) List() SubscribedPulses {
	client := &http.Client{}
	req, _ := http.NewRequest(get, baseURL + subscriptionsURLPath, nil)
	req.Header.Set("X-OTX-API-KEY", fmt.Sprintf("%s", APIKey))
	response, _ := client.Do(req)

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	pulses := SubscribedPulses{}
	json.Unmarshal(contents, &(pulses.data))
	return pulses
}