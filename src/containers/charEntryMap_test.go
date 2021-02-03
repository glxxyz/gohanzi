package containers_test

import (
	"github.com/glxxyz/gohanzi/containers"
	"testing"
)

func TestCharEntryMap_AddContains(t *testing.T) {
	sut := containers.CharEntryMap{}
	sut['A'] = EntryA
	sut['B'] = EntryB
	assert(t, sut.Contains('A'), "contains A")
	assert(t, sut.Contains('B'), "contains B")
	assert(t, !sut.Contains('C'), "doesn't contain C")
	sut['C'] = EntryC
	assert(t, sut.Contains('C'), "contains C")
}

func TestCharEntryMap_Difference(t *testing.T) {
	AB := containers.CharEntryMap{}
	AB['A'] = EntryA
	AB['B'] = EntryB
	BC := containers.CharEntryMap{}
	BC['B'] = EntryB
	BC['C'] = EntryC
	result := AB.Difference(BC)
	expected := containers.CharEntryMap{}
	expected['A'] = EntryA
	assertEquals(t, expected, result)
}

func TestCharEntryMap_Intersection(t *testing.T) {
	AB := containers.CharEntryMap{}
	AB['A'] = EntryA
	AB['B'] = EntryB
	BC := containers.CharEntryMap{}
	BC['B'] = EntryB
	BC['C'] = EntryC
	result := AB.Intersection(BC)
	expected := containers.CharEntryMap{}
	expected['B'] = EntryB
	assertEquals(t, expected, result)
}

func TestCharEntryMap_SubSetOf(t *testing.T) {
	ABC := containers.CharEntryMap{}
	ABC['A'] = EntryA
	ABC['B'] = EntryB
	ABC['C'] = EntryC
	BC := containers.CharEntryMap{}
	BC['B'] = EntryB
	BC['C'] = EntryC
	ABZ := containers.CharEntryMap{}
	ABZ['A'] = EntryA
	ABZ['B'] = EntryB
	ABZ['Z'] = EntryZ
	assert(t, ABC.SubsetOf(ABC), "ABC is a subset of ABC")
	assert(t, BC.SubsetOf(ABC), "BC is a subset of ABC")
	assert(t, !ABC.SubsetOf(BC), "ABC is not a subset of BC")
	assert(t, !ABC.SubsetOf(ABZ), "ABC is not a subset of ABZ")
}

func TestCharEntryMap_Union(t *testing.T) {
	AB := containers.CharEntryMap{}
	AB['A'] = EntryA
	AB['B'] = EntryB
	BC := containers.CharEntryMap{}
	BC['B'] = EntryB
	BC['C'] = EntryC
	result := AB.Union(BC)
	assertEquals(t, 2, len(AB))
	assertEquals(t, 2, len(BC))
	expected := containers.CharEntryMap{}
	expected['A'] = EntryA
	expected['B'] = EntryB
	expected['C'] = EntryC
	assertEquals(t, expected, result)
}

func TestCharEntryMap_AddAll(t *testing.T) {
	sut := containers.CharEntryMap{}
	sut['A'] = EntryA
	sut['B'] = EntryB
	BC := containers.CharEntryMap{}
	BC['B'] = EntryB
	BC['C'] = EntryC
	sut.AddAll(BC)
	expected := containers.CharEntryMap{}
	expected['A'] = EntryA
	expected['B'] = EntryB
	expected['C'] = EntryC
	assertEquals(t, expected, sut)
	assertEquals(t, 2, len(BC))
}
