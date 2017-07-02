xorshift [![Build Status](https://travis-ci.org/lazybeaver/xorshift.svg?branch=master)](https://travis-ci.org/lazybeaver/xorshift)
========
A golang implementation of the xorshift family of pseudorandom number generators.

These generators are simultaneously high quality, fast, cache-line friendly, simple to implement, have a large period and pass the [TestU01](http://simul.iro.umontreal.ca/testu01/tu01.html) statistical tests.

xorshift1024* and xorshift128+ seem to be strictly better than the Mersenne Twister and WELL family of PRNGs in every dimension. For more details, see the [xorshift webpage](http://xorshift.di.unimi.it/#shootout).

References
==========

- http://www.jstatsoft.org/v08/i14/paper
- http://vigna.di.unimi.it/ftp/papers/xorshift.pdf
- http://vigna.di.unimi.it/ftp/papers/xorshiftplus.pdf

Example
========

```go
  package main

  import (
    "fmt"
    "time"  
    "github.com/lazybeaver/xorshift"
  )

  func main() {
    seed := uint64(time.Now().Nanosecond())

    xor64s := xorshift.NewXorShift64Star(seed)
    fmt.Printf("xorshift64*    : %x\n", xor64s.Next())

    xor128p := xorshift.NewXorShift128Plus(seed)
    fmt.Printf("xorshift128+   : %x\n", xor128p.Next())

    xor1024s := xorshift.NewXorShift1024Star(seed)
    fmt.Printf("xorshift1024s  : %x\n", xor1024s.Next())
  }
```

Benchmarks
==========

```
  $ go test -bench=.
  PASS
  BenchmarkXorShift1024Star	500000000	         7.20 ns/op
  BenchmarkXorShift128Plus	500000000	         3.83 ns/op
  BenchmarkXorShift64Star	500000000	         6.83 ns/op
  ok  	github.com/lazybeaver/xorshift	10.748s
```
