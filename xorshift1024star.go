package xorshift

type XorShift1024Star struct {
	state [16]uint64
	index uint32
}

func (x *XorShift1024Star) Next() uint64 {
	var s0, s1 uint64
	s0 = x.state[x.index]
	x.index = (x.index + 1) & 15
	s1 = x.state[x.index]
	s1 ^= (s1 << 31)
	s1 ^= (s1 >> 11)
	s0 ^= (s0 >> 30)
	x.state[x.index] = s0 ^ s1
	return x.state[x.index] * 1181783497276652981
}
