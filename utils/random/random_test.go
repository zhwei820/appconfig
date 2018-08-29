package random

import (
	"testing"
	"math/rand"
)

func TestRandString(t *testing.T) {
	s := RandString(10)
	println(s)
}

func BenchmarkRandString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandString(10)
	}
}

func BenchmarkRandIntr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rand.Int()
	}
}

func BenchmarkRandInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandInt(1, 11)
	}
}
