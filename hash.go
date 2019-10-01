package phi2_hash

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lphi2_hash
// #include "phi2.h"
import "C"

import (
	"unsafe"
)

func GetPowHash(hash []byte) []byte {
	powhash := make([]byte, 32)
	C.phi2_hash((*C.char)(unsafe.Pointer(&hash[0])), (*C.char)(unsafe.Pointer(&powhash[0])))
	return powhash
}
