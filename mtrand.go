package npoker

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"github.com/seehuhn/mt19937"
	"os"
)

const (
	N64   = 312
	NByte = 2496
)

type Rand struct {
	mt19937.MT19937
}

func NewRand() *Rand {
	return &Rand{*mt19937.New()}
}

/*
	from "math/rand"
*/

func (r *Rand) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	return int(r.Int63n(int64(n)))
}

func (r *Rand) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { //n is power of two, can mask
		return r.Int63() & (n - 1)
	}

	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Int63()
	for v > max {
		v = r.Int63()
	}
	return v % n
}

/*
func (r *Rand) Int63() int64 {
	return r
}
*/

func GetSeedFromURAND() []byte {
	b := make([]byte, NByte)
	_, err := rand.Reader.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func GetSeedFromRAND() []byte {
	b := make([]byte, NByte)
	f, err := os.Open("/dev/random")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	for i := 0; i < NByte; {
		n, err := f.Read(b)
		if err != nil {
			panic(err)
		}
		i += n
	}
	return b
}

func Dump(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Loads(s string) []byte {
	b, _ := base64.StdEncoding.DecodeString(s)
	return b
}

func (r *Rand) SeedFromBytes(b []byte) {
	/* wrapper for SeedFromSlice, for ease of use */
	if len(b) != NByte {
		panic("Seed bytes must be 2496 bytes long")
	}
	u := make([]uint64, N64)
	for i := 0; i < N64; i++ {
		u[i], _ = binary.Uvarint(b[i*8 : (i+1)*8])
	}
	r.SeedFromSlice(u)
}
