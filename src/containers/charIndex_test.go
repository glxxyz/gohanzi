package containers_test

import (
	"github.com/glxxyz/gohanzi/containers"
	"testing"
)

func TestCharIndex_Add(t *testing.T) {
	sut := containers.CharIndex{}
	sut.Add('1', EntryA)
	sut.Add('1', EntryB)
	sut.Add('2', EntryC)
	assert(t, sut['1'].Contains(EntryA), "1 has EntryA")
	assert(t, sut['1'].Contains(EntryB), "1 has EntryB")
	assert(t, sut['2'].Contains(EntryC), "2 has EntryC")
	assert(t, !sut['2'].Contains(EntryB), "2 doesn't have EntryB")
}

func TestCharIndex_AddAll(t *testing.T) {
	sut := containers.CharIndex{}
	sut.Add('1', EntryA)
	other := containers.CharIndex{}
	other.Add('1', EntryB)
	other.Add('2', EntryC)
	assert(t, sut['1'].Contains(EntryA), "1 has EntryA")
	assert(t, !sut['1'].Contains(EntryB), "1 doesn't have EntryB")
	assert(t, !sut['2'].Contains(EntryC), "2 doesn't have EntryC")
	assert(t, !sut['2'].Contains(EntryB), "2 doesn't have EntryB")
	sut.AddAll(other)
	assert(t, sut['1'].Contains(EntryA), "1 has EntryA")
	assert(t, sut['1'].Contains(EntryB), "1 has EntryB")
	assert(t, sut['2'].Contains(EntryC), "2 has EntryC")
	assert(t, !sut['2'].Contains(EntryB), "2 doesn't have EntryB")
}

