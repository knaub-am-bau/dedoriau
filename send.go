///////////////////////////////////////////////////////////////
// Send a zero-value transaction with a "docker-id" message //
// This first draft is based on the code of the IOTA guide //
////////////////////////////////////////////////////////////

package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/bundle"
    "github.com/iotaledger/iota.go/converter"
    "github.com/iotaledger/iota.go/trinary"
    "fmt"

)
// devnet
var node = "https://nodes.devnet.thetangle.org"

// Define a seed and an address.
// These do not need to belong to anyone or have IOTA tokens.
// They must only contain a maximum of 81 trytes
// or 90 trytes with a valid checksum
const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
const address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")

// Define a message to send.
// This message must include only ASCII characters.
// This part must be changed to a variable with the ID
var data = "{'message' : 'Hello world'}"

const minimumWeightMagnitude = 9
const depth = 3

func main() {
    // compose a new API instance, we provide no PoW function so this uses remote PoW
    api, err := ComposeAPI(HTTPClientSettings{URI: node})
    must(err)

    // Convert the message to trytes
    message, err := converter.ASCIIToTrytes(data)
    must(err)

    // Define a zero-value transaction object
    // that sends the message to the address
    transfers := bundle.Transfers{
        {
            Address: address,
            Value: 0,
            Message: message,
        },
    }
    // Use the default options
    prepTransferOpts := PrepareTransfersOptions{}

    trytes, err := api.PrepareTransfers(seed, transfers, prepTransferOpts)
    must(err)

    // Create a bundle from the `transfers` object
    // and send the transaction to the node
    myBundle, err := api.SendTrytes(trytes, depth, minimumWeightMagnitude)
    must(err)

    fmt.Println("Bundle hash: " + myBundle[0].Bundle)
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
