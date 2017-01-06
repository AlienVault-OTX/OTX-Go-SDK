// Copyright 2015 The go-otxapi AUTHORS. All rights reserved.
//
// Use of this source code is governed by an Apache
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AlienVault-Labs/OTX-Go-SDK/src/otxapi"
)

func main() {
	one := flag.Bool("one", false, "Only print one pulse, then exit")
	flag.Parse()

	// Create a new client, searching the environment for an API key.
	client, err := otxapi.NewClient(otxapi.APIKeyFromEnv())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Get your user information.
	user, err := client.User.Details()
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not get user details:", err)
		os.Exit(2)
	}
	fmt.Println(user)

	// Get the first page of pulses we are subscribed to, and pretty-print
	// them.
	plist, err := client.Pulses.List(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not get first page of pulses:", err)
		os.Exit(2)
	}
	if *one {
		fmt.Println(plist.Pulses[0])
		return
	}
	for _, pulse := range plist.Pulses {
		fmt.Println(pulse)
	}

	// Get the rest of the pages of pulses we are subscribed to, and
	// pretty-print them.
	for {
		opts, err := plist.NextPageOptions()
		if err != nil && err == otxapi.ErrNoPage {
			// We are on the last page.
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}

		plist, err = client.Pulses.List(opts)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error getting next page of results:", err)
			os.Exit(2)
		}
		for _, pulse := range plist.Pulses {
			fmt.Println(pulse)
		}
	}
}
