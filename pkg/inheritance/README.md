# Inheritance

Go does not have inheritence, it has struct embedding. Most of the time, that's good enough but there are times where traditional inheritance can be useful.

This package uses the classic shape hierarchy since it's familiar but whenever you have some "base" type that would benefit from calling methods on "child" types, you might benefit from this pattern. In the shape example, the "base" type is Shape and we need to call the "child" methods `Name` and `Area` on the "child" types: Circle, Rectangle, Square, and BadTriangle.

This pattern has two distinct problems, both demonstrated in this example. Neither is particularly bad in practice once you know about them.

1. If you forget to implement a method in the "child" for which there is no "parent" method, you'll get a stack overflow error at runtime. 
2. You must construct the child types correctly. If not, you'll get nil pointer errors trying to call the inheritable methods.

A more practical example of where this might be useful is in a SQL storage implementation where most of the functionality was generic, but some required specialization for a particular database. In this case, all the generic code can be implemented on the "base" type, which can call out to specialized versions where needed.

