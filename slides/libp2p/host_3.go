package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-peerstore"
)

func main() {

	// START OMIT
	// Simple hosts with one address each
	p1, _ := libp2p.New(context.Background(), libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
		libp2p.DisableRelay())
	p2, _ := libp2p.New(context.Background(), libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
		libp2p.DisableRelay())

	// Peers and peerstores
	fmt.Printf("1: %v %v %v\n", p1.ID(), p1.Addrs(), p1.Peerstore().Peers())
	fmt.Printf("2: %v %v %v\n\n", p2.ID(), p2.Addrs(), p2.Peerstore().Peers())

	// Add addresses to peerstores
	p1.Peerstore().AddAddrs(p2.ID(), p2.Addrs(), peerstore.PermanentAddrTTL)
	p2.Peerstore().AddAddrs(p1.ID(), p1.Addrs(), peerstore.PermanentAddrTTL)

	// Peerstores and addresses
	fmt.Printf("1: %v %v\n", p1.Peerstore().PeerInfo(p1.ID()), p1.Peerstore().PeerInfo(p2.ID()))
	fmt.Printf("2: %v %v\n", p2.Peerstore().PeerInfo(p1.ID()), p2.Peerstore().PeerInfo(p2.ID()))

	//
	// END OMIT
}
