Go Patterns
=====

On the surface, Go's system of structs and interfaces can seem primitive compared to familiar OO concepts. However, a surprising amount of functionality can be achieved by following certain patterns that may not be obvious at first glance. This repository provides guides to these patterns and describes the applicability, strengths, and weaknesses of each.

Every pattern has tradeoffs and the reader is the responsibility to balance those tradeoffs when selecting to use and not use patterns in this repository. By including the pattern, I neither encourage nor discourage it's usage; instead, I strive to be exhaustive to serve the maintainer who may encounters "bad" patterns in the wild.

- [x] [Enums](pkg/enums/README.md)
- [x] [Inheritance](pkg/inheritance/README.md)
- [x] [Marker Traits](pkg/markers/README.md) 
- [x] [Tagged Unions](pkg/markers/markers.go) (aka sum type, variant, enum)
- [x] [Protected](pkg/protected/README.md)
- [x] [Private](pkg/private/README.md) struct access.
- [ ] Writers
- [ ] Cofiguration
- [ ] Default Parameters