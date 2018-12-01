package main

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-crypto"
)

func main() {

	// START OMIT
	// Set own keypair
	privKey, _, _ := crypto.GenerateEd25519Key(rand.Reader)

	// Host with some options
	host, _ := libp2p.New(context.Background(),
		libp2p.Identity(privKey),
		libp2p.DisableRelay(),
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))

	// ID
	fmt.Printf("ID: %s (%s)\n", host.ID().Pretty(), host.ID())

	// Addresses
	for i, addr := range host.Addrs() {
		fmt.Printf("addr %d: %v\n", i, addr)
	}

	// Multiplexer protocols
	for i, protoc := range host.Mux().Protocols() {
		fmt.Printf("protoc %d: %v\n", i, protoc)
	}
	// END OMIT
}
