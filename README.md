# iabconsent
[![Build Status][ci]](https://travis-ci.com/LiveRamp/iabconsent)

[ci]: https://travis-ci.com/LiveRamp/iabconsent.svg?branch=master "Build Status"

[![Go Report Card][report]](https://goreportcard.com/report/github.com/LiveRamp/iabconsent)

[report]: https://goreportcard.com/badge/github.com/LiveRamp/iabconsent "Go Report Card"

A Golang implementation of the IAB Consent String 1.1 Spec as well as the IAB Transparency and Consent String v2.0.

To install:
```
go get -v github.com/LiveRamp/iabconsent
```

This package defines a two structs (`ParsedConsent` and `V2ParsedConsent`) which contain all of the fields of the IAB 
v1.1 Consent String and the IAB Transparency and Consent String v2.0 respectively.

Each spec has their own parsing function (`ParseV1` and `ParseV2`). If the caller is unsure of which version a consent
string is, they can use `ParseVersion` to receive a `StringVersion` to determine which parsing function is appropriate.

Example use:
```go
package main

import "github.com/LiveRamp/iabconsent"

func main() {
    var consent = "COvzTO5OvzTO5BRAAAENAPCoALIAADgAAAAAAewAwABAAlAB6ABBFAAA"

    var c, err = iabconsent.ParseVersion(consent)
    if err != nil {
        panic(err)
    }
    
    switch c {
    case iabconsent.V1:
        var v1, err = iabconsent.ParseV1(consent)
        // Use v1 consent.
    case iabconsent.V2:
        var v2, err = iabconsent.ParseV2(consent)
        // Use v2 consent.
    default:
        panic("unknown version")
    }
}
```

The function `Parse(s string)` is deprecated, and should no longer be used.
