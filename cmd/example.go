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
