package containers_test

import (
	"github.com/glxxyz/gohanzi/containers"
	"testing"
)

func TestStringIndex_Add(t *testing.T) {
	sut := containers.StringIndex{}
	sut.Add("word1", EntryA)
	sut.Add("word1", EntryB)
	sut.Add("word2", EntryC)
	assert(t, sut["word1"].Contains(EntryA), "word1 has EntryA")
	assert(t, sut["word1"].Contains(EntryB), "word1 has EntryB")
	assert(t, sut["word2"].Contains(EntryC), "word2 has EntryC")
	assert(t, !sut["word2"].Contains(EntryB), "word2 doesn't have EntryB")
}

func TestStringIndex_AddAll(t *testing.T) {
	sut := containers.StringIndex{}
	sut.Add("word1", EntryA)
	other := containers.StringIndex{}
	other.Add("word1", EntryB)
	other.Add("word2", EntryC)
	assert(t, sut["word1"].Contains(EntryA), "word1 has EntryA")
	assert(t, !sut["word1"].Contains(EntryB), "word1 doesn't have EntryB")
	assert(t, !sut["word2"].Contains(EntryC), "word2 doesn't have EntryC")
	assert(t, !sut["word2"].Contains(EntryB), "word2 doesn't have EntryB")
	sut.AddAll(other)
	assert(t, sut["word1"].Contains(EntryA), "word1 has EntryA")
	assert(t, sut["word1"].Contains(EntryB), "word1 has EntryB")
	assert(t, sut["word2"].Contains(EntryC), "word2 has EntryC")
	assert(t, !sut["word2"].Contains(EntryB), "word2 doesn't have EntryB")
}
