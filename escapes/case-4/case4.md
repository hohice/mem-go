```shell
$ go build -gcflags '-m -l' main.go
# command-line-arguments
./main.go:9:18: leaking param: u to result ~r1 level=0
./main.go:15:18: &User literal does not escape
```
leaking param 的表述，它说明了变量 u 是一个泄露参数。结合代码可得知其传给 GetUserInfo 方法后，没有做任何引用之类的涉及变量的动作，直接就把这个变量返回出去了。

因此这个变量实际上并没有逃逸，它的作用域还在 main() 之中，所以分配在栈上。