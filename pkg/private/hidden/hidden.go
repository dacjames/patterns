package hidden

type daxRequestFailure struct {
	codes []int
}

type daxTransactionCanceledFailure struct {
	daxRequestFailure
	cancellationReasonCodes []string
	cancellationReasonMsgs  []string
	cancellationReasonItems []byte
}

func MakeHidden() *daxTransactionCanceledFailure {
	return &daxTransactionCanceledFailure{
		daxRequestFailure: daxRequestFailure{
			codes: []int{400, 404},
		},
		cancellationReasonCodes: []string{"client error", "not found"},
		cancellationReasonMsgs:  []string{"malformed request parameter foobar", ""},
		cancellationReasonItems: []byte("{}"),
	}
}
