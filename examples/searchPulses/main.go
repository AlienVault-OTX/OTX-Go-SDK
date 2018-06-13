// Copyright 2015 The go-otxapi AUTHORS. All rights reserved.
//
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"os"
	"otxapi"
)

func main() {
	os.Setenv("X_OTX_API_KEY", "myotxapikey")
	client := otxapi.NewClient(nil)

	searchString := flag.String("q", "", "Search string")
	flag.Parse()

	opt := &otxapi.ListOptions{Query: *searchString}
	pulseSearchResponse, _, _ := client.SearchPulses.Search(opt)
	fmt.Printf("Number of results: %d\n\n", *pulseSearchResponse.Count)
	for _, result := range pulseSearchResponse.Results {
		fmt.Printf("%s - %s\n", *result.ID, *result.Name)
		for _, indicator := range result.Indicators {
			fmt.Printf("    %s = %s\n", *indicator.Type, *indicator.Indicator)
		}
	}
}
