///////////////////////////////////////////////////////////////
// send a transaction with the docker hash                  //
// This first draft is based on the code of the IOTA guide //
// not funktional at the moment                           //
///////////////////////////////////////////////////////////

package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/trinary"
    "github.com/iotaledger/iota.go/bundle"
    "fmt"

)

const node = "https://nodes.devnet.thetangle.org"

// Replace this seed with the one that owns the address you used to get free test tokens
const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")

const minimumWeightMagnitude = 9
const depth = 3

func main() {
    // Connect to a node
    api, err := ComposeAPI(HTTPClientSettings{URI: node})
    must(err)

    // Define an address to which to send IOTA tokens
    address := trinary.Trytes("ZLGVEQ9JUZZWCZXLWVNTHBDX9G9KZTJP9VEERIIFHY9SIQKYBVAHIMLHXPQVE9IXFDDXNHQINXJDRPFDXNYVAPLZAW")

    // Define an input transaction object
    // that sends 1 i to your new address
    transfers := bundle.Transfers{
          {
              Address: address,
              Value: 1,
          },
      }

    fmt.Println("Sending 1 i to " + address);

    trytes, err := api.PrepareTransfers(seed, transfers, PrepareTransfersOptions{})
    must(err)

    myBundle, err := api.SendTrytes(trytes, depth, minimumWeightMagnitude)
    must(err)

    fmt.Println(myBundle)
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
