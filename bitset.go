package main

import "math/bits"

type bitset struct {
	words []uint64
}

func newBitset() *bitset {
	return &bitset{words: make([]uint64, 1)}
}

func (b *bitset) Set(bit uint, value bool) {
	n := bit / 64
	if n >= uint(len(b.words)) {
		b.words = append(b.words, make([]uint64, n-uint(len(b.words))+1)...)
	}
	mask := uint64(1) << (bit % 64)
	if value {
		b.words[n] |= mask
	} else {
		b.words[n] = b.words[n] | mask ^ mask
	}
}

func (b *bitset) Test(bit uint) bool {
	n := bit / 64
	if n >= uint(len(b.words)) {
		return false
	}
	return b.words[n]&(1<<(bit%64)) != 0
}

func (b *bitset) OrWith(a *bitset) {
	if len(a.words) > len(b.words) {
		b.words = append(b.words, make([]uint64, len(a.words)-len(b.words))...)
	}
	for i, word := range a.words {
		b.words[i] |= word
	}
}

func (b *bitset) BitCount() int {
	return Sum(Map(b.words, func(w uint64) int {
		return bits.OnesCount64(w)
	}))
}
