[![GoDoc](https://pkg.go.dev/badge/github.com/ianmclinden/qmk-go)](https://pkg.go.dev/github.com/ianmclinden/qmk-go)


# qmk-go

Go client for VIA-enabled QMK keyboards.

---

## Compiling

qmk-go depends on [karalabe/hid](https://github.com/karalabe/hid) and is therfore subject to the same compilation restrictions:

> Cross compiling to a different platform or architecture entails disabling CGO by default in Go, causing device enumeration hid.Enumerate() to yield no results.
> 
> To cross compile a functional version of this library, you'll need to enable CGO during cross compilation via CGO_ENABLED=1 and you'll need to install and set a cross compilation enabled C toolkit via CC=your-cross-gcc.

## Example

```Go
package main

import (
    "fmt"

    "github.com/ianmclinden/qmk-go"
)

const (
    vendorId  = 0x4b54 // USB VendorID
    productId = 0x2323 // USB ProductID
    serial    = ""     // Optional
)

func main() {
    // Create a new keyboard connection
    client, err := qmk.NewClient(vendorId, productId, serial)
    if err != nil {
        fmt.Errorf(err)
    }

    version, err := client.GetProtocolVersion()
    if (err != nil) {
        fmt.Errorf(err)
    }

    fmt.Printf("Connected to keyboard with VIA version 0x%04x\n", version)
}
```

## License
Because of the license requirements from [karalabe/hid](https://github.com/karalabe/hid) this library is licensed under GNU LGPL 2.1 or later on Linux and 3-clause BSD on other platforms.