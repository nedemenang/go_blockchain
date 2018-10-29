package blockchain

import (
	. "github.com/nedemenang/go_blockchain/blocks"
)


type Blockchain struct {
	Blocks []*Block
}

var BlockChain *Blockchain

func(bc *Blockchain) AddBlock (data BookCheckout) {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]

	block := CreateBlock(prevBlock, data)
	if validBlock(block, prevBlock) {
		bc.Blocks = append(bc.Blocks, block)
	}
}


func GenesisBlock() *Block {
	return CreateBlock(&Block{}, BookCheckout{IsGenesis: true})
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}


func validBlock(block, prevBlock *Block) bool {
	
	if prevBlock.Hash != block.PrevHash {
		return false
	}

	if !block.ValidateHash(block.Hash) {
		return false
	}

	if prevBlock.Pos + 1 != block.Pos{
		return false 
	}

	return true
}