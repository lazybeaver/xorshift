package xorshift

// A generic xorshift PRNG interface
type XorShift interface {
	Next() uint64
}

// Return a xorshift64* generator with a period of 2^64 - 1
// Use this only as a seed generator for other PRNGs
// Reference: http://vigna.di.unimi.it/ftp/papers/xorshift.pdf
func NewXorShift64Star(seed uint64) XorShift {
	return &XorShift64Star{state: seed}
}

// Return a xorshift128+ generator with a period of 2^128 - 1
// Reference: http://vigna.di.unimi.it/ftp/papers/xorshiftplus.pdf
func NewXorShift128Plus(seed uint64) XorShift {
	xor128p := &XorShift128Plus{}
	for i, seedgen := 0, NewXorShift64Star(seed); i < len(xor128p.state); i++ {
		xor128p.state[i] = seedgen.Next()
	}
	return xor128p
}

// Return a xorshift1024* generator with a period of 2^1024 - 1
// Reference: http://vigna.di.unimi.it/ftp/papers/xorshift.pdf
func NewXorShift1024Star(seed uint64) XorShift {
	xor1024s := &XorShift1024Star{}
	for i, seedgen := 0, NewXorShift64Star(seed); i < len(xor1024s.state); i++ {
		xor1024s.state[i] = seedgen.Next()
	}
	return xor1024s
}
