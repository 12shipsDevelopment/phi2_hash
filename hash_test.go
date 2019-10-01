package phi2_hash

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestPhi2Hash(t *testing.T) {
	hash := "00000060748aa994d1e375f029af8c1134f86724fc47a92c4ce9729cfa5127c10eb60bb8820a3a0c4ea474d9330cfd563b5a69d25ed3e37bd63ddd5b22e6ac3c811cc780c29fea5c3fd9511b00196b10"
	hash_bin, _ := hex.DecodeString(hash)
	powhash := GetPowHash(hash_bin)
	powhash_hex := hex.EncodeToString(powhash)
	fmt.Println(powhash_hex)
	if powhash_hex != "77a19463753c27887c5697b47118719f4af6fba0647eddde71a938e7b3dd0d48" {
		t.Error("Test x16r powhash failed")
		return
	}
}
