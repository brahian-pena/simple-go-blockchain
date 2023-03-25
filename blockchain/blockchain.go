package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

type Block struct {
	current_hash  string
	previous_hash string
	proof         int
	index         int
	timestamp     int64
	data          string
}

var blockchain []*Block

func InitBlockchain() {
	println("Initializing blockchain")
	CreateBlock(1, "0", "{}")
	println("Successfully initialized blockchain")
}

func CreateBlock(proof int, previous_hash string, data string) (b *Block) {

	block := &Block{
		previous_hash: previous_hash,
		index:         len(blockchain) + 1,
		timestamp:     time.Now().Unix(),
		proof:         proof,
		data:          data,
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

	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return hash
}

func ProofOfWork(previous_proof int) {
	new_proof := 1

	check_proof := false

	for !check_proof {
		hash_operation := CreateHash(float64ToByte(math.Pow(float64(new_proof), 2) - math.Pow(float64(previous_proof), 2)))
	}
}

func float64ToByte(f float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
