package main

import (
	"fmt"

	mt "github.com/txaty/go-merkletree"
	"golang.org/x/crypto/sha3"
)

// first define a data structure with Serialize method to be used as data block
type testData struct {
	data []byte
}

func (t *testData) Serialize() ([]byte, error) {
	return t.data, nil
}

func keccak256Hash(data []byte) ([]byte, error) {

	h := sha3.NewLegacyKeccak256()
	h.Write(data)
	hash := h.Sum(nil)
	//fmt.Printf("%x\n", hash) // b10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6

	return hash, nil
}

func main() {
	blocks := make([]mt.DataBlock, 0)

	for i := 0; i < 100; i++ {
		block := &testData{
			data: []byte("0x5b2495F3D183628Faf891b64A28D7A392D8b3759"),
		}
		blocks = append(blocks, block)
	}
	// VIDA VIDA VIDA

	// the first argument is config, if it is nil, then default config is adopted
	c := &mt.Config{
		HashFunc: keccak256Hash,
	}

	tree, err := mt.New(c, blocks)
	handleError(err)
	// get proofs
	proofs := tree.Proofs
	// or you can also verify the proofs without the tree but with Merkle root
	// obtain the Merkle root
	rootHash := tree.Root
	for i := 0; i < len(blocks); i++ {
		// if hashFunc is nil, use SHA256 by default
		ok, err := mt.Verify(blocks[i], proofs[i], rootHash, keccak256Hash)
		handleError(err)
		fmt.Println(ok)
		fmt.Println(i)

		for i := 0; i < len(proofs[i].Siblings); i++ {
			fmt.Printf("%x\n", proofs[i].Siblings[i])
		}

	}
	fmt.Printf("rootHash: %x\n", rootHash)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
