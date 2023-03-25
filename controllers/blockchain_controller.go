package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/brahian-pena/simple-go-blockchain/blockchain"
)

func MineBlock(c *fiber.Ctx) error {
	previous_block := blockchain.GetLastBlock()

	proof := blockchain.ProofOfWork(previous_block.Proof)

	block := blockchain.CreateBlock(proof, blockchain.HashBlock(previous_block), string(c.Body()))

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Block added succesfully",
		"status":  fiber.StatusCreated,
		"result":  block,
	})
}

func GetBlockChain(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"result": blockchain.GetBlockChain(),
	})
}

func ValidateBlockChain(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"result": blockchain.IsBlockchainValid(),
	})
}
