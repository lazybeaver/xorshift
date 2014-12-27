package xorshift

import (
	"testing"
)

func TestXorShift64Star(t *testing.T) {
	w := &XorShift64Star{state: uint64(1)}
	expected := [5]uint64{0x47e4ce4b896cdd1d, 0xabcfa6a8e079651d, 0xb9d10d8feb731f57, 0x4db418a0bb1b019d, 0xe6199b04d5aa600}
	actual := [5]uint64{0, 0, 0, 0, 0}
	for i := 0; i < len(expected); i++ {
		actual[i] = w.Next()
		if actual[i] != expected[i] {
			t.Errorf("Mismatch for iteration %d: Actual=%x vs Expected=%x", i, actual[i], expected[i])
		}
	}
}

func BenchmarkXorShift64Star(b *testing.B) {
	w := NewXorShift64Star(0x1122334455667788)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.Next()
	}
}
