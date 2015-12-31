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

    var PulsesStruct otxapi.SubscribedPulses
	p := PulsesStruct.List()
	fmt.Println(p)

}
