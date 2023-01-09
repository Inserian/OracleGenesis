package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// The Genesis block
type GenesisBlock struct {
	Timestamp     int64
	Data          string
	PrevBlockHash string
	Hash          string
}

// The Oracle block
type OracleBlock struct {
	Timestamp     int64
	Data          string
	PrevBlockHash string
	Hash          string
}

// The Blockchain struct
type Blockchain struct {
	GenesisBlock GenesisBlock
	Database     []OracleBlock
}

// The smart contract function for the Genesis block
func (b *GenesisBlock) SetData(data string) {
	b.Data = data
	b.Timestamp = time.Now().Unix()
	b.PrevBlockHash = ""
	b.Hash = b.calculateHash()
}

// The smart contract function for the Oracle block
func (b *OracleBlock) SetData(data string) {
	b.Data = data
	b.Timestamp = time.Now().Unix()
	b.PrevBlockHash = b.calculateHash()
	b.Hash = b.calculateHash()
}

// The function to calculate the hash of a block
func (b *GenesisBlock) calculateHash() string {
	record := string(b.Timestamp) + b.Data + b.PrevBlockHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *OracleBlock) calculateHash() string {
	record := string(b.Timestamp) + b.Data + b.PrevBlockHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// The function to create a new Genesis block
func CreateGenesisBlock(data string) GenesisBlock {
	block := GenesisBlock{}
	block.SetData(data)
	return block
}

// The function to create a new Oracle block
func CreateOracleBlock(data string, prevBlock GenesisBlock) OracleBlock {
	block := OracleBlock{}
	block.SetData(data)
	block.PrevBlockHash = prevBlock.Hash
	return block
}

func main() {
	// Create the Genesis block
	genesisBlock := CreateGenesisBlock("Genesis Block")
	fmt.Println("Genesis Block:")
	fmt.Println(genesisBlock)
	
	// Create the Oracle block
	databaseBlock := CreateOracleBlock("Oracle Block", genesisBlock)
	fmt.Println("Oracle Block:")
	fmt.Println(databaseBlock)
}
