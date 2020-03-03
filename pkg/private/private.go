package main

import (
	"fmt"
	"unsafe"

	"github.com/dacjames/patterns/pkg/private/hidden"
)

type DaxRequestFailure struct {
	Codes []int
}

type DaxTransactionCanceledFailure struct {
	DaxRequestFailure
	CancellationReasonCodes []string
	CancellationReasonMsgs  []string
	CancellationReasonItems []byte
}

func main() {
	h := hidden.MakeHidden()
	// h is already an pointer, take it's address if not.
	pub := (*DaxTransactionCanceledFailure)(unsafe.Pointer(h))

	fmt.Printf("Public: %#v", pub)
}
