package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {

	// START OMIT
	// Simple host with all the default settings
	host, _ := libp2p.New(context.Background())

	// ID
	fmt.Printf("ID: %s (%s)\n", host.ID().Pretty(), host.ID())

	// Addresses
	for i, addr := range host.Addrs() {
		fmt.Printf("addr %d: %v\n", i, addr)
	}

	// Peerstore
	fmt.Printf("Peerstore %T %v\n", host.Peerstore(), host.Peerstore().Peers())

	// Connection Manager
	fmt.Printf("ConnManager() %T %v\n", host.ConnManager(), host.ConnManager())

	// Multiplexer protocols
	for i, protoc := range host.Mux().Protocols() {
		fmt.Printf("protoc %d: %v\n", i, protoc)
	}

	// Network
	fmt.Printf("network: %v\n", host.Network())
	// END OMIT
}
