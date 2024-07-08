# Benchmarking Demo

This repo covers benchmarking for the top four solutions to the Leetcode problem [350. Intersection of Two Arrays II](https://leetcode.com/problems/intersection-of-two-arrays-ii/description/?envType=daily-question&envId=2024-07-02).

The original test cases were small (array length < 10), so these benchmarks demonstrate how the different solutions behave across various input sizes:
- 1
- 10
- 100
- 1,000
- 10,000
- 100,000
- 1,000,000

The slices are pregenerated before each round with a ~20% array intersection, and max array size has been limited to ensure that swap isn't hit (at least on my machine).

> Note: I have not checked these solutions for correctness and they are just examples for demonstration purposes.

## Running 

### Generate benchmarks for all four solutions

```shell
$ mkdir results
$ go test -benchmem -bench=. -timeout=1h -count=10 -args 1 > results/1.txt
$ go test -benchmem -bench=. -timeout=1h -count=10 -args 2 > results/2.txt
$ go test -benchmem -bench=. -timeout=1h -count=10 -args 3 > results/3.txt
$ go test -benchmem -bench=. -timeout=1h -count=10 -args 4 > results/4.txt
```

### Generate a single benchmark

```shell
$ go test -benchmem -bench=. -timeout=30m -count=10 -args 1 
```
Where "1" is the solution algorithm, either "1", "2", "3", or "4".

> Note: pregenerated sample benchmarks are found in the `samples` directory.

## Comparing

[`benchstat`](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) can calculate differences with a confidence interval, so let's show that:
```shell
$ go install golang.org/x/perf/cmd/benchstat@latest
$ cd results
$ benchstat *.txt
```

> `benchstat` can only compare the benchmarks if the test names are the same, which is the reason we have been running the same benchmark function, and just passing a command line arg to switch between the different "backing" algorithms.

## Future tweaks

Further parameters can be adjusted for these benchmarks:
- amount of array intersection 
- length difference between the arrays