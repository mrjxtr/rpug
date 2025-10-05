package generator

import (
	"encoding/hex"
	"testing"
)

func TestSeedGenerator(t *testing.T) {
	const seedLen = 32

	seed, err := generateSeed()
	if err != nil {
		t.Fatalf("expected no erro, got: %v", err)
	}

	if len(seed) != seedLen {
		t.Errorf("expected seed length %d, got: %d", seedLen, len(seed))
	}

	if _, err := hex.DecodeString(seed); err != nil {
		t.Errorf("seed is not valid hex: %v", err)
	}
}

func TestRNGDeterminism(t *testing.T) {
	const seed = "8959bcbac47d82c434fd8f154dab3e04"

	rngA := newRNGfromSeed(seed)
	rngB := newRNGfromSeed(seed)

	for range 5 {
		n1 := rngA.Intn(1000)
		n2 := rngB.Intn(1000)
		if n1 != n2 {
			t.Errorf(
				"expected deterministic output for same seed but got %d != %d",
				n1,
				n2,
			)
		}
	}
}

func TestRNGUniqueness(t *testing.T) {
	const n = 1000
	rngA := newRNGfromSeed("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	rngB := newRNGfromSeed("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	rngC := newRNGfromSeed("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaac")
	rngD := newRNGfromSeed("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbd")

	// Generate a few numbers and check they differ
	same := true
	i := 0
	for range 5 {
		i++
		a := rngA.Intn(n)
		b := rngB.Intn(n)
		c := rngC.Intn(n)
		d := rngD.Intn(n)

		t.Logf("[%d] A:%d B:%d C:%d D:%d", i, a, b, c, d)

		if a != b || a != c || a != d {
			same = false
		}
	}
	if same {
		t.Errorf(
			"expected different RNG outputs for different seeds, but got identical results",
		)
	}
}
