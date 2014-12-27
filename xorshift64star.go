package xorshift

type XorShift64Star struct {
	state uint64
}

func (x *XorShift64Star) Next() uint64 {
	x.state ^= (x.state >> 12)
	x.state ^= (x.state << 25)
	x.state ^= (x.state >> 27)
	return x.state * 2685821657736338717
}
