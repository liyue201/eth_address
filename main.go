package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethersphere/bee/pkg/crypto"
)

func genAddress() {
	k, _ := crypto.GenerateSecp256k1Key()
	d1 := crypto.EncodeSecp256k1PrivateKey(k)
	singer := crypto.NewDefaultSigner(k)
	addr, _ := singer.EthereumAddress()

	fmt.Printf("%v %v\n", hex.EncodeToString(d1), addr.String())
}

func main() {
	for i := 0; i < 10; i++ {
		genAddress()
	}
}
