# Readme

## Questions and Answers

### Diff between Elem and Indirect?

- https://prnt.sc/26dauyc
-

If a reflect.Value is a pointer, then v.Elem() is equivalent to reflect.Indirect(v). If it is not a pointer, then they
are not equivalent:

If the value is an interface then reflect.Indirect(v) will return the same value, while v.Elem() will return the
contained dynamic value. If the value is something else, then v.Elem() will panic. The reflect.Indirect helper is
intended for cases where you want to accept either a particular type, or a pointer to that type. One example is the
database/sql conversion routines: by using reflect.Indirect, it can use the same code paths to handle the various types
and pointers to those types.

## Links

* [go - golang - Elem Vs Indirect in the reflect package - Stack Overflow](https://stackoverflow.com/questions/24318389/golang-elem-vs-indirect-in-the-reflect-package#:~:text=Elem%20returns%20the%20value%20that,Value%20if%20v%20is%20nil.&text=Indirect%20returns%20the%20value%20that%20v%20points%20to.)
* [reflect.Indirect() Function in Golang with Examples - GeeksforGeeks](https://www.geeksforgeeks.org/reflect-indirect-function-in-golang-with-examples/)
