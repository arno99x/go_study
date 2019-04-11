package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  //区块链创建的时间戳
	Data          []byte //区块链包含的数据
	PrevBlockHash []byte //前一个区块链hash值
	Hash          []byte //区块链自身哈希值，用于校验数据有效性
	Nonce         int    //
}

/**
创建新的block
打上当前的时间戳，输入block data，再加入上一个块的hash值
*/
func NewBlock(data string, prevBlockHash []byte) *Block {
	newBlock := &Block{Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlockHash, Hash: []byte{}}
	pow := NewProofofWork(newBlock)
	nonce, hash := pow.Run()

	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce
	return newBlock
}

/**
为block设置hash值
为了block被篡改，hash值中包括：当前block创建时的时间戳，上一个块的hash值，当前块的data
*/
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	header := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(header)
	block.Hash = hash[:]
}

/**
创建 blockchain中的第一个节点，被叫做genesis block，也叫创世纪块
*/
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
