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

import (
  "fmt"

  "github.com/LiveRamp/iabconsent"
)

func main() {
    var v1, err = iabconsent.ParseV1("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", v1)

    var v2 *iabconsent.V2ParsedConsent
    v2, err = iabconsent.ParseV2("real-value-goes-here")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", v2)

    c, err = iabconsent.ParseVersion("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", c)
}
```

The function `Parse(s string)` is deprecated, and should no longer be used.
