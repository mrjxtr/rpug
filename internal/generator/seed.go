package generator

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"log/slog"
	mathrand "math/rand"
)

func generateSeed() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	seed := hex.EncodeToString(bytes)

	slog.Info("Seed generated", "seed", seed)
	return seed, nil
}

func seedToInt(seed string) int64 {
	hash := md5.Sum([]byte(seed))
	return int64(binary.BigEndian.Uint64(hash[:8]))
}

func newRNGfromSeed(seed string) *mathrand.Rand {
	source := mathrand.NewSource(seedToInt(seed))
	return mathrand.New(source)
}
