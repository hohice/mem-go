```shell
$ go build -gcflags '-m -l' main.go
# command-line-arguments
./main.go:4:12: new(string) does not escape
```
该对象分配到栈上了。很核心的一点就是它有没有被作用域之外所引用，而这里作用域仍然保留在 main 中，因此它没有发生逃逸