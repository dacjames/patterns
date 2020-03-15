# Marker Traits

Marker traits can be used in place of `interface{}` types to wrap a restricted set of structs. This pattern is useful in two distinct scenarios:

- Adding type information to a function signature. 
- Mimicing a *tagged union* (aka `enum`, aka `sealed trait`) of struct types.

## Function Signatures

It is common for Go packages to define custom type aliases for empty interface types to serve as self documentation. For example, the [crypto]() package defines two such type aliases:

```go
package crypto

type PrivateKey interface{}

type PublicKey interface{}
```

The subpackages of `crypto` define concrete type aliases with matching names:

```go
package ed25519

type PrivateKey []byte

type PubliceKey []byte
```

The aliasing creates nice function signatures, like the `Sign` function:

```go
package ed25519
func Sign(privateKey PrivateKey, message []byte) []byte { return nil }

```

So long as we're only working with the `ed25519` package, the type aliases protect us using the wrong type. However, we can run into problems when trying to generalize over different PublicKeys.

```go
package example
import (
    "crypto"
    "crypto/ed25519"
)

func Encrypt(plaintext []byte, publicKey crypto.PublicKey) ([]byte, error) { return nil, nil }

func Example() {
    pub, priv, _ := ed25519.GenerateKey(nil)
    good, _ := Encrypt([]byte{"hello"}, pub)
    bad, _ := Encrypt([]byte{"hello"}, priv) // oops! we're trying use a private key, when a public key is required!
}
```

Depending on the implementation of the Encrypt function and how good our error handling is, the impact of this issue can range from a minor annoyance to a serious bug. IN this case, a silent failure is a practical possibility because `ed25519.PublicKey` and `ed25519.PrivateKey` are both aliases for the same concrete type, `[]byte`). To avoid this problem, we need to constrain the `crypto.PublicKey` type such that it can contain any of the concrete PublicKey types but only PublicKey types. That is, we need a way to *mark* the concrete `ed25519.PublicKey` as a `crypto.PublicKey`.

The solution is almost emberassingly simple: define a trivial method on `crypto.PrivateKey` and implement it for every concrete `PrivateKey` type.

```go
package crypto2

type PrivateKey interface{
    MarkPrivateKey()
}

type PublicKey interface{
    MarkPublicKey()
}
```
```go
package ed25519

type PrivateKey []byte
func (PrivateKey) MarkPrivateKey() {}

type PubliceKey []byte
func (PublicKey) MarkPublicKey() {}
```

The name of the method is irrelevant but I conventionally use `MarkWhatever` or `markWhatever`. With this small modification, we can catch the bug in our example at compile time.

```go
package example
import (
    "crypto2"
    "crypto2/ed25519"
)

func Example() {
    pub, priv, _ := ed25519.GenerateKey(nil)
    good, _ := Encrypt([]byte{"hello"}, pub)
    bad, _ := Encrypt([]byte{"hello"}, priv) // error: ed25519.PrivateKey does not implement crypto2.PublicKey.
}
```

### Is this worth it?

Marker traits can prevent bugs but they require a certain amount of boilerplate to implement and use. This boilerplate can leak into the documented interface of the function, as it does in our example. The marker method must be public whenever the types implementing it live in a different package than the marker interface.

As with all these patterns, making the tradeoff is up to the reader. Having experiences these type of bugs in the past, I personally find the additional protection worthwhile.

## Tagged Unions

Many programming languages, particular those that encourage functional programming pattern, provide tagged unions in the language directly, usally under a different name to avoid confusion with the less powerful *untagged* unions available in C or with union types. In scala2, any class can be used as a tagged union by adding it to a `sealed trait` or `sealed abstract class`; in Rust, they are spelled `enum`. We can simulate tagged unions in Go using marker traits.

A representative implementation can be found in [`markers.go`](./markers.go).





