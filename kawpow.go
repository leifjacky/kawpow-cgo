package kawpow

/*
void kawpow_hash(char *, char *, const char*, const char*, const char*);
#cgo LDFLAGS: -L. -Llib -lkawpow
#include <stdlib.h>
*/
import "C"
import (
	"encoding/hex"
	"fmt"
	"sync"
	"unsafe"
)

var mu sync.Mutex

func KawpowHash(headerHash, nonce string, height int64) ([]byte, []byte) {
	mu.Lock()
	defer mu.Unlock()
	heightStr := C.CString(fmt.Sprintf("%d", height))
	defer C.free(unsafe.Pointer(heightStr))

	headerHashStr := C.CString(headerHash)
	defer C.free(unsafe.Pointer(headerHashStr))

	nonceStr := C.CString(nonce)
	defer C.free(unsafe.Pointer(nonceStr))

	hash := make([]byte, 64)
	mixHash := make([]byte, 64)
	C.kawpow_hash((*C.char)(unsafe.Pointer(&hash[0])), (*C.char)(unsafe.Pointer(&mixHash[0])), headerHashStr, nonceStr, heightStr)

	hashT, _ := hex.DecodeString(string(hash))
	mixHashT, _ := hex.DecodeString(string(mixHash))
	return hashT, mixHashT
}
