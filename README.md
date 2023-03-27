
# This project keeps a few benchmarks written for studying Golang

## Which is faster: String keys or struct keys?

When using maps Go does allow us to use structs as keys when we have multiple values.

So the question arrived, would it be better to just concatenate several values in the form of a
big string and use it as the unique key of the map, or would it be faster to have a struct
that contains all these values?

This benchmarks discusses this possibility:

```bash
$ make
go test -bench=. -benchtime=5s
goos: darwin
goarch: amd64
pkg: github.com/vingarcia/go-benchmarks
cpu: VirtualApple @ 2.50GHz
BenchmarkMapKeys/using_strings-10         	28784778	       207.6 ns/op
BenchmarkMapKeys/using_strings_optimized-10 50145543	       126.2 ns/op
BenchmarkMapKeys/using_structs-10           87705032	       67.55 ns/op
PASS
ok  	github.com/vingarcia/go-benchmarks	19.321s
```
