package xorshift

import (
	"testing"
)

func TestXorShift128Plus(t *testing.T) {
	w := &XorShift128Plus{state: [2]uint64{1, 2}}
	expected := [5]uint64{0x800045, 0x2000104, 0x4000020010c3, 0xc00002103045, 0x1000801c450c4}
	actual := [5]uint64{0, 0, 0, 0, 0}
	for i := 0; i < len(expected); i++ {
		actual[i] = w.Next()
		if actual[i] != expected[i] {
			t.Errorf("Mismatch for iteration %d: Actual=%x vs Expected=%x", i, actual[i], expected[i])
		}
	}
}

func BenchmarkXorShift128Plus(b *testing.B) {
	w := NewXorShift128Plus(0x1122334455667788)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.Next()
	}
}
