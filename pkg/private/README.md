# Private Members

Somtimes, you really want to access a private struct in some go package. A common example is accessing an error the package authors did not think you needed.

To get around this limiation, you can make a copy of the private struct and use `unsafe.Pointer` to cast away the privateness. 

This should be safe so long as the copied struct is exactly the same as the private struct.