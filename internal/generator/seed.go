package generator

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	mathrand "math/rand"
)

// generateSeed returns a 16-byte hex string sourced from crypto/rand.
// Use this seed to reproduce the same "random" data later.
func generateSeed() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	seed := hex.EncodeToString(bytes)

	return seed, nil
}

// newRNGfromSeed creates a math/rand RNG that's deterministic for the seed.
// Same seed in, same random sequence out.
func newRNGfromSeed(seed string) *mathrand.Rand {
	source := mathrand.NewSource(seedToInt(seed))
	return mathrand.New(source)
}

// seedToInt converts a hex seed into a stable int64 using an md5 hash.
// Small seed changes produce a very different output value.
func seedToInt(seed string) int64 {
	hash := md5.Sum([]byte(seed))
	return int64(binary.BigEndian.Uint64(hash[:8]))
}
