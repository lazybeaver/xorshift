package xorshift

import (
	"testing"
)

func TestXorShift1024Star(t *testing.T) {
	w := &XorShift1024Star{state: [16]uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}
	expected := [5]uint64{0xc0562e31b467f91f, 0x92b6fabadaff6d4, 0x6a37d6c71bffb6a, 0xd534ffc84bb7e231, 0x61cf9e3dc667e6c7}
	actual := [5]uint64{0, 0, 0, 0, 0}
	for i := 0; i < len(expected); i++ {
		actual[i] = w.Next()
		if actual[i] != expected[i] {
			t.Errorf("Mismatch for iteration %d: Actual=%x vs Expected=%x", i, actual[i], expected[i])
		}
	}
}

func BenchmarkXorShift1024Star(b *testing.B) {
	w := NewXorShift1024Star(0x1122334455667788)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.Next()
	}
}
