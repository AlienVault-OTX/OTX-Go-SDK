// Copyright 2015 The go-otxapi AUTHORS. All rights reserved.
//
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"otxapi"
	"log"
)

func main() {

	client := otxapi.NewClient(nil)
	// get user details, easy way to validate api key
	user_detail, _, err := client.UserDetail.Get()
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n", otxapi.Stringify(user_detail))
	}

	valid_user := *(user_detail.UserId) > 0
	if !valid_user {
		log.Fatal("Please set your otx api key as the environment variable X_OTX_API_KEY.")
	}
	fmt.Println("Valid User: ", valid_user)
	id_string := "56cdb0a04637f275671672f3"
	pulse_detail, _, err := client.PulseDetail.Get(id_string)

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n", otxapi.Stringify(pulse_detail))
	}
	opt := &otxapi.ListOptions{Page: 1, PerPage: 4}
	pulse_list, _, err := client.ThreatIntel.List(opt)

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n", otxapi.Stringify(pulse_list))
	}
}
