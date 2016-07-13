package npoker

import (
	"fmt"
	"testing"
)

func TestLoadSeed(t *testing.T) {
	b := GetSeedFromURAND()
	r := NewRand()
	r.SeedFromBytes(b)
	fmt.Println(r.Int63())

}

func BenchmarkGetSeedFromURAND(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetSeedFromURAND()
	}
}

func BenchmarkIntn(b *testing.B) {
	r := NewRand()
	r.SeedFromBytes(GetSeedFromURAND())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Int63()
	}
}
