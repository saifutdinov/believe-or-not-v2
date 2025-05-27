package cards

import (
	"crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

func newRNG() *mrand.Rand {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic("seeding RNG: " + err.Error())
	}
	seed := int64(binary.LittleEndian.Uint64(b[:]))
	return mrand.New(mrand.NewSource(seed))
}

func Shuffle(deck []*Card) []*Card {
	rng := newRNG()
	shuffled := make([]*Card, len(deck))
	copy(shuffled, deck)
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	return shuffled
}
