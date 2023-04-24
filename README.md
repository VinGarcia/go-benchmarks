
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
BenchmarkMapKeys/using_strings-10         	27575566	       215.7 ns/op
BenchmarkMapKeys/using_strings_optimized-10         	57606266	       106.6 ns/op
BenchmarkMapKeys/using_structs-10                   	86174137	        69.67 ns/op
PASS
ok  	github.com/vingarcia/go-benchmarks	24.823s
```

## What is the fasted way to work with a Matrix in Golang?

When creating matrixes I've seen some people using a single big array to represent
matrixes supposedly for performance reasons. So I decided to put that to proof.

This benchmark compares an implementation of a NxN square matrix using three
different implementations:

1. A slice of slices where each sub-slice points the a contiguous array in memory which might improve the performance
2. A non-optimized slice of slices where each sub-slice might be allocated in a different way than the previous one.
3. A single big array where the position i,j is calculated by the following formula: `i*200 + j`

```bash
$ make
go test -bench=. -benchtime=5s
goos: darwin
goarch: amd64
pkg: github.com/vingarcia/go-benchmarks
cpu: VirtualApple @ 2.50GHz
BenchmarkMatrixIndexing/slice_of_slices-10                  	1000000000	         1.766 ns/op
BenchmarkMatrixIndexing/slice_of_slices_single_allocation-10	1000000000	         1.766 ns/op
BenchmarkMatrixIndexing/single_slice-10                     	1000000000	         1.600 ns/op
PASS
ok  	github.com/vingarcia/go-benchmarks	24.823s
```
