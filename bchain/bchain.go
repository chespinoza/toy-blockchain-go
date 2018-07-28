package bchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block ...
type Block struct {
	Index     int
	TimeStamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// calculateHash returns hash containing the data
func calculateHash(block Block) string {
	record := string(block.Index) + block.TimeStamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// generateBlock ...
func GenerateBlock(oldblock Block, BPM int) (Block, error) {
	var newBlock Block
	newBlock.Index = oldblock.Index + 1
	newBlock.TimeStamp = time.Now().String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldblock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil

}

// validateBlock checks the new block which is wanted to be added
func ValidateBlock(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// replaceChain() compare chains and write the larger over the shorter
func ReplaceChain(newBlocks []Block, currBlocks *[]Block) {
	if len(newBlocks) > len(*currBlocks) {
		currBlocks = &newBlocks
	}
}
