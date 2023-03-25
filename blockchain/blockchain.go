package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/brahian-pena/simple-go-blockchain/utils"
)

type Block struct {
	PreviousHash string `json:"previous_hash"`
	Proof        int    `json:"proof"`
	Index        int    `json:"index"`
	Timestamp    int64  `json:"timestamp"`
	Data         string `json:"data"`
}

var blockchain []*Block

const ACCEPTANCE_CRITERIA = "0000"

func InitBlockchain() {
	println("Initializing blockchain")
	CreateBlock(1, "0", "{}")
	println("Successfully initialized blockchain")
}

func CreateBlock(proof int, previous_hash string, data string) (b *Block) {

	block := &Block{
		PreviousHash: previous_hash,
		Index:        len(blockchain) + 1,
		Timestamp:    time.Now().Unix(),
		Proof:        proof,
		Data:         data,
	}

	blockchain = append(blockchain, block)

	return block
}

func GetLastBlock() (b *Block) {
	return blockchain[len(blockchain)-1]
}

func CreateHash(bytes_to_hash []byte) string {
	hasher := sha256.New()

	hasher.Write(bytes_to_hash)

	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}

func ProofOfWork(previous_proof int) int {
	new_proof := 1

	check_proof := false

	for !check_proof {
		hash_operation := CreateHash(utils.Float64ToByte(math.Pow(float64(new_proof), 2) - math.Pow(float64(previous_proof), 2)))

		if strings.HasPrefix(hash_operation, ACCEPTANCE_CRITERIA) {
			check_proof = true
		} else {
			new_proof++
		}
	}

	return new_proof
}

func HashBlock(block *Block) string {
	encoded_block, err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var hash = CreateHash(encoded_block)

	return hash
}

func IsBlockchainValid() bool {
	previous_block := blockchain[0]

	block_index := 1

	for block_index < len(blockchain) {
		current_block := blockchain[block_index]

		if current_block.PreviousHash != HashBlock(previous_block) {
			return false
		}

		hash := CreateHash(utils.Float64ToByte(math.Pow(float64(current_block.Proof), 2) - math.Pow(float64(previous_block.Proof), 2)))

		if !strings.HasPrefix(hash, ACCEPTANCE_CRITERIA) {
			return false
		}

		previous_block = current_block

		block_index++
	}

	return true
}

func GetBlockChain() []*Block {
	return blockchain
}
