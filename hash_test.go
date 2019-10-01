package phi2_hash

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestPhi2Hash(t *testing.T) {
	hash := "00000060748aa994d1e375f029af8c1134f86724fc47a92c4ce9729cfa5127c10eb60bb8820a3a0c4ea474d9330cfd563b5a69d25ed3e37bd63ddd5b22e6ac3c811cc780c29fea5c3fd9511b00196b1094fee56ea92723d9bf0a097dac2f933cec3be22538a5a7d9ab42a4d490acd161365f0a5b4d44b49054f7b8980e94890caedd7ebb00c92089b93f3fbaa55f8f73"
	hash_bin, _ := hex.DecodeString(hash)
	powhash := GetPowHash(hash_bin)
	powhash_hex := hex.EncodeToString(powhash)
	if powhash_hex != "97e8cc4b019de5a610ad5198ba112fa78e6793df5818f95f7b6b440000000000" {
		t.Error("Test phi2 powhash failed")
		return
	}
}
