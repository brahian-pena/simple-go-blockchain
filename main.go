package main

import (
	"github.com/brahian-pena/simple-go-blockchain/blockchain"
	"github.com/brahian-pena/simple-go-blockchain/server"
)

func main() {
	println("Running GO HomeMade Blockchain Server")

	blockchain.InitBlockchain()

	server.StartListening()
}
