/////////////////////////////////////////////////////////
// Address needed to push the Docker ID to the tangle //
///////////////////////////////////////////////////////

package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/trinary"
    "fmt"
)
// a seed must be created first
// this part needs to be automated in the future
var seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
var endpoint = "https://nodes.devnet.thetangle.org"

func main() {
    // compose a new API instance
    api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
    must(err)

    // GetNewAddress retrieves the first unspent from address through IRI
    addresses, err := api.GetNewAddress(seed, GetNewAddressOptions{})
    must(err)

    fmt.Println("\nYour new address: ", addresses[0])
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
