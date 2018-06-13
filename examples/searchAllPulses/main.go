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
	os.Setenv("X_OTX_API_KEY", "myapikey")
	client := otxapi.NewClient(nil)

	searchString := flag.String("q", "", "Search string")
	flag.Parse()

	opt := &otxapi.ListOptions{Query: *searchString}

	pulseCh := make(chan otxapi.Pulse, 10)
	responseCh := make(chan otxapi.Response, 10)
	doneCh := make(chan bool)

	go func() {
		err := SearchAllPulses(client, opt, pulseCh, responseCh)
		if err != nil {
			fmt.Println(err)
      close(pulseCh)
      close(responseCh)
		}
	}()

	go func(doneCh chan bool) {
		for result := range pulseCh {
			fmt.Printf("%v : %v\n", *result.ID, *result.Name)
		}
		doneCh <- true
	}(doneCh)

	<-doneCh
}

func SearchAllPulses(c *otxapi.Client, opt *otxapi.ListOptions, pulseCh chan otxapi.Pulse, responseCh chan otxapi.Response) error {
	if opt.PerPage == 0 {
		opt.PerPage = 5
	}
	opt.Page = 1
	pulseSearchResponse, resp, err := c.SearchPulses.Search(opt)
	if err != nil {
		return err
	}

	for _, result := range pulseSearchResponse.Results {
		pulseCh <- result
	}
	responseCh <- resp

	var pages int

	if *pulseSearchResponse.Count > 0 {
		pages = *pulseSearchResponse.Count / opt.PerPage
	} else {
		pages = 1
	}

	for opt.Page = 2; opt.Page <= pages; opt.Page++ {
		pulseSearchResponse, resp, err = c.SearchPulses.Search(opt)
		if err != nil {
			return err
		}

		for _, result := range pulseSearchResponse.Results {
			pulseCh <- result
		}
		responseCh <- resp
	}

	close(pulseCh)
	close(responseCh)

	return err
}
