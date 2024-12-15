package main

import "testing"

func TestBitset(t *testing.T) {
	b := newBitset()
	CheckBitsetSingleBit(t, b, 0)
	CheckBitsetSingleBit(t, b, 1)
	CheckBitsetSingleBit(t, b, 10)
	CheckBitsetSingleBit(t, b, 63)
	CheckBitsetSingleBit(t, b, 64)
	CheckBitsetSingleBit(t, b, 65)
}

func TestBitset_OrWith(t *testing.T) {
	a := newBitset()
	b := newBitset()
	b.Set(1, true)
	b.Set(65, true)
	a.Set(2, true)
	a.OrWith(b)
	AssertBitNotSet(t, a, 0)
	AssertBitSet(t, a, 1)
	AssertBitSet(t, a, 2)
	AssertBitNotSet(t, a, 64)
	AssertBitSet(t, a, 65)
}

func TestBitset_BitCount(t *testing.T) {
	b := newBitset()
	AssertEquals(t, 0, b.BitCount())
	b.Set(0, true)
	AssertEquals(t, 1, b.BitCount())
	b.Set(65, true)
	AssertEquals(t, 2, b.BitCount())
	b.Set(0, false)
	AssertEquals(t, 1, b.BitCount())
	b.Set(65, false)
	AssertEquals(t, 0, b.BitCount())
}

func CheckBitsetSingleBit(t *testing.T, b *bitset, bit uint) {
	AssertBitNotSet(t, b, bit)
	b.Set(bit, true)
	AssertBitSet(t, b, bit)
	b.Set(bit, false)
	AssertBitNotSet(t, b, bit)
}

func AssertBitSet(t *testing.T, b *bitset, bit uint) {
	if !b.Test(bit) {
		t.Errorf("bitset.Test(%d) should return true", bit)
	}
}

func AssertBitNotSet(t *testing.T, b *bitset, bit uint) {
	if b.Test(bit) {
		t.Errorf("bitset.Test(%d) should return false", bit)
	}
}
