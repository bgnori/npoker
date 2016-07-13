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
