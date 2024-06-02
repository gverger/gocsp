package variables

import (
	"log"

	"github.com/bits-and-blooms/bitset"
)

type BitSetDomain struct {
	Values *bitset.BitSet
}

// Size implements Domain.
func (b *BitSetDomain) Size() int {
	return int(b.Values.Count())
}

func newBitsetDomain(domSize int) *BitSetDomain {
	return &BitSetDomain{
		Values: bitset.New(uint(domSize)).SetAll(),
	}
}

// Clone implements Domain.
func (b *BitSetDomain) Clone() Domain {
	return &BitSetDomain{Values: b.Values.Clone()}
}

// Empty implements Domain.
func (b *BitSetDomain) Empty() bool {
	return b.Values.None()
}

// Fix implements Domain.
func (b *BitSetDomain) Fix(value int) bool {
	if !b.Values.Test(uint(value)) {
		return false
	}

	b.Values.ClearAll()
	b.Values.Set(uint(value))
	return true
}

// Fixed implements Domain.
func (b *BitSetDomain) Fixed() bool {
	return b.Values.Count() == 1
}

// Min implements Domain.
func (b BitSetDomain) Min() int {
	set, ok := b.Values.NextSet(0)
	if !ok {
		log.Fatal("Min of empty set")
	}
	return int(set)
}

// Remove implements Domain.
func (b *BitSetDomain) Remove(value int) bool {
	if !b.Values.Test(uint(value)) {
		return false
	}

	b.Values.Clear(uint(value))
	return true
}

var _ Domain = &BitSetDomain{}
