## Protected

This package demonstrates shared a semi-private implementation between an "owner" and a "friend," while preventing an "enemy" package from inadvertently referencing the implementation. This simulates the concept of protected methods found in OO languages.

The basic idea is to put the methods in a public interface, but implement them for a private struct. You then store a private reference to an instance of that struct in the owner and the friend uses an unsafe cast to gain access to the private variable.

The "enemy" here can utilize the same cast as the friend, so this only protects accidental dependence.

