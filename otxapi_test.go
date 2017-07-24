package otxapi

import (
	"fmt"
	"log"
	"net/http"
)

func ExampleAPIKey() {
	client, err := NewClient(APIKey("..."))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client)
}

func ExampleAPIKeyFromEnv_default() {
	client, err := NewClient(APIKeyFromEnv())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client)
}

func ExampleAPIKeyFromEnv_customVarNames() {
	client, err := NewClient(APIKeyFromEnv(
		"ALIENVAULT_OTX_KEY",
		"OTX_KEY",
		"OTX_API_KEY",
	))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client)
}

func ExampleHTTPClient() {
	customClient := &http.Client{}
	client, err := NewClient(HTTPClient(customClient))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client)
}

func ExampleUserAgent() {
	client, err := NewClient(UserAgent("custom-otx-api-client/13.37"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client)
}
