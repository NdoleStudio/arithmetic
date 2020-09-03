arithmetic
==========

This package provides basic arithmetic functions without the need for type casting when working with go.

Benchmarks
==========

This is a benchmark for computing the same functions using this library and doing it natively using type casting. 
I'm running this tests on a late 2018 a Macbook Pro with 16 GB RAM and a 2.3 GHz Quad-Core Intel Core i5 processor.


#### DivideIntsReturnFloat

```go
DivideIntsReturnFloat(1, 3)
// vs
float63(1) / float64(3)
```

```text
BenchmarkDivideIntsReturnFloat-8                 1000000000               0.572 ns/op
BenchmarkDivideIntsReturnFloatNative-8           1000000000               0.553 ns/op
```

#### MinInt

```go
MinInt(31, 42)
// vs
Math.Min(float63(31), 42)
```

```text
BenchmarkMinInt-8                                1000000000               0.545 ns/op
BenchmarkMathMin-8                               437190435                2.75 ns/op
```

#### MinInts

```go
MinInts(42,90, 30)
```

```text
BenchmarkMinInts-8                               256053702                4.82 ns/op
```

#### MaxInt

```go
MaxInt(31, 42)
// vs
Math.Max(float63(31), 42)
```

```text
BenchmarkMaxInt-8                                1000000000               0.576 ns/op
BenchmarkMathMax-8                               404535982                2.93 ns/op
```

#### MaxInts

```go
MaxInts(42,90, 30)
```
```text
BenchmarkMaxInts-8                               275597671                4.32 ns/op
```