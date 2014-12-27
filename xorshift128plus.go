package xorshift

type XorShift128Plus struct {
	state [2]uint64
}

func (x *XorShift128Plus) Next() uint64 {
	var s0, s1 uint64
	s1 = x.state[0]
	s0 = x.state[1]
	x.state[0] = s0
	s1 ^= (s1 << 23)
	x.state[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26))
	return (x.state[0] + x.state[1])
}
