package friend

import (
	"unsafe"

	"github.com/dacjames/patterns/pkg/protected/owner"
)

type extractor struct {
	Impl owner.Protected
}

var ownerProtected = (*extractor)(unsafe.Pointer(&owner.Protector)).Impl

func SayHelloTwice() {
	ownerProtected.SayHello()
}
