package otxapi

import (
	"fmt"
	"log"
)

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
