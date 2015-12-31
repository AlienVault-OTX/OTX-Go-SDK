// Copyright 2015 The go-otxapi AUTHORS. All rights reserved.
//
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"otxapi"
)

func main() {
//	client := otxapi.NewClient(nil)
//	fmt.Println(client)
//	fmt.Println("Recently updated repositories owned by user willnorris:")
////
//	otxapi.SubscriptionList()
//	opt := &otxapi.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
//	repos, _, err := client.Repositories.List("willnorris", opt)
//	if err != nil {
//		fmt.Printf("error: %v\n\n", err)
//	} else {
//		fmt.Printf("%v\n\n", github.Stringify(repos))
//	}

    var PulsesStruct otxapi.SubscribedPulses
	p := PulsesStruct.List()
	fmt.Println(p)

}
