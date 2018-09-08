# Enums

The enums package documents how to build progressively more advandeced enum types.

See [enums.go](/pkg/enums/enums.go) for documented enum definitions.
See [enums_example_test.go](/pkg/enums/enums_example_test.go) for testable enum usage example.


Three "levels" of enums are defined.

|pattern|pros|cons|example|
|---|---|---|---|
|iota enums|super simple, backwards compatible|no type safety, no convenience methods|Cardinal|
|newtype enums|type safety, convenience methods, minal overhead|literal gotchas, no protection against invalid enums|Direction|
|interface enums|prevents creation of invalid enums|overhead of interface type, verbosity/complexity|Icon|
