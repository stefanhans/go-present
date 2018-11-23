package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {

	// START_1 OMIT
	// Simple host with all the default settings
	h, _ := libp2p.New(context.Background())

	// ID
	fmt.Printf("ID: %s (%s)\n", h.ID().Pretty(), h.ID())

	// Addresses
	for i, addr := range h.Addrs() {
		fmt.Printf("addr %d: %v\n", i, addr)
	}

	// Peerstore
	fmt.Printf("Peerstore %T %v\n", h.Peerstore(), h.Peerstore().Peers())

	// Connection Manager
	fmt.Printf("ConnManager() %T %v\n", h.ConnManager(), h.ConnManager())

	// Multiplexer protocols
	for i, protoc := range h.Mux().Protocols() {
		fmt.Printf("protoc %d: %v\n", i, protoc)
	}

	// Network
	fmt.Printf("network: %v\n", h.Network())
	// END_1 OMIT

	// START_2 OMIT
	// Simple host with all the default settings
	h, _ = libp2p.New(context.Background())

	// ID
	fmt.Printf("ID: %s (%s)\n", h.ID().Pretty(), h.ID())

	// Addresses
	for i, addr := range h.Addrs() {
		fmt.Printf("addr %d: %v\n", i, addr)
	}

	// Peerstore
	fmt.Printf("Peerstore %T %v\n", h.Peerstore(), h.Peerstore().Peers())

	// Connection Manager
	fmt.Printf("ConnManager() %T %v\n", h.ConnManager(), h.ConnManager())

	// Multiplexer protocols
	for i, protoc := range h.Mux().Protocols() {
		fmt.Printf("protoc %d: %v\n", i, protoc)
	}

	// Network
	fmt.Printf("network: %v\n", h.Network())
	// END_2 OMIT
}
