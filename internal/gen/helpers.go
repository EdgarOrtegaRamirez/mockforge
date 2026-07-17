package gen

import (
	"math/rand/v2"
)

func randForSeed(seed int64) *rand.Rand {
	if seed == 0 {
		seed = 0
	}
	return rand.New(rand.NewPCG(uint64(seed), uint64(seed^0x6A09E667F3BCC908)))
}