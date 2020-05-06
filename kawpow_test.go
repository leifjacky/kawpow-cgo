package kawpow

import (
	"fmt"
	"testing"
)

func TestKawpowHash(t *testing.T) {
	headerHash := "acd14e518a5e8c8f200e7233e627a63ec32f6f0a47a499a6e4b5f530e153153d"
	nonce := "763df482051ed3dd"
	height := int64(274070)

	hash, mixHash := KawpowHash(headerHash, nonce, height)
	fmt.Printf("hash: %064x, mixHash: %064x\n", hash, mixHash)
}
