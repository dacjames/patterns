package owner

type protected struct{}

func (_ protected) SayHello() string {
	return "Hello, world!"
}

type Protected interface {
	SayHello() string
}

type protector struct {
	impl Protected
}

var Protector = &protector{impl: protected{}}

func SayHelloThrice() {
	Protector.impl.SayHello()
	Protector.impl.SayHello()
	Protector.impl.SayHello()
}
