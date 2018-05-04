# iabconsent
A Golang implementation of the IAB Consent String 1.1 Spec

To install:
```
go get -v github.com/LiveRamp/iabconsent
```

This package defines a struct (`ParsedConsent`) which contains all of the fields of the IAB Consent String. The function `Parse(s string)` accepts the Base64 Raw URL Encoded cookie string and returns a `ParsedConsent` with all relevent fields populated.

Example use:
```go
package main

import (
  "fmt"

  "github.com/LiveRamp/iabconsent"
)

func main() {
  var pc, err = iabconsent.Parse("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
  if err != nil {
    panic(err)
  }
  fmt.Printf("%+v\n", pc)
}
```

