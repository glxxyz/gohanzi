package containers_test

import (
	"github.com/glxxyz/gohanzi/containers"
	"testing"
)

func TestStringEntryMap_AddContains(t *testing.T) {
	sut := containers.StringEntryMap{}
	sut["A"] = EntryA
	sut["B"] = EntryB
	assert(t, sut.Contains("A"), "contains A")
	assert(t, sut.Contains("B"), "contains B")
	assert(t, !sut.Contains("C"), "doesn't contain C")
	sut["C"] = EntryC
	assert(t, sut.Contains("C"), "contains C")
}

func TestStringEntryMap_Difference(t *testing.T) {
	AB := containers.StringEntryMap{}
	AB["A"] = EntryA
	AB["B"] = EntryB
	BC := containers.StringEntryMap{}
	BC["B"] = EntryB
	BC["C"] = EntryC
	result := AB.Difference(BC)
	expected := containers.StringEntryMap{}
	expected["A"] = EntryA
	assertEquals(t, expected, result)
}

func TestStringEntryMap_Intersection(t *testing.T) {
	AB := containers.StringEntryMap{}
	AB["A"] = EntryA
	AB["B"] = EntryB
	BC := containers.StringEntryMap{}
	BC["B"] = EntryB
	BC["C"] = EntryC
	result := AB.Intersection(BC)
	expected := containers.StringEntryMap{}
	expected["B"] = EntryB
	assertEquals(t, expected, result)
}

func TestStringEntryMap_SubSetOf(t *testing.T) {
	ABC := containers.StringEntryMap{}
	ABC["A"] = EntryA
	ABC["B"] = EntryB
	ABC["C"] = EntryC
	BC := containers.StringEntryMap{}
	BC["B"] = EntryB
	BC["C"] = EntryC
	ABZ := containers.StringEntryMap{}
	ABZ["A"] = EntryA
	ABZ["B"] = EntryB
	ABZ["Z"] = EntryZ
	assert(t, ABC.SubsetOf(ABC), "ABC is a subset of ABC")
	assert(t, BC.SubsetOf(ABC), "BC is a subset of ABC")
	assert(t, !ABC.SubsetOf(BC), "ABC is not a subset of BC")
	assert(t, !ABC.SubsetOf(ABZ), "ABC is not a subset of ABZ")
}

func TestStringEntryMap_Union(t *testing.T) {
	AB := containers.StringEntryMap{}
	AB["A"] = EntryA
	AB["B"] = EntryB
	BC := containers.StringEntryMap{}
	BC["B"] = EntryB
	BC["C"] = EntryC
	result := AB.Union(BC)
	assertEquals(t, 2, len(AB))
	assertEquals(t, 2, len(BC))
	expected := containers.StringEntryMap{}
	expected["A"] = EntryA
	expected["B"] = EntryB
	expected["C"] = EntryC
	assertEquals(t, expected, result)
}

func TestStringEntryMap_AddAll(t *testing.T) {
	sut := containers.StringEntryMap{}
	sut["A"] = EntryA
	sut["B"] = EntryB
	BC := containers.StringEntryMap{}
	BC["B"] = EntryB
	BC["C"] = EntryC
	sut.AddAll(BC)
	expected := containers.StringEntryMap{}
	expected["A"] = EntryA
	expected["B"] = EntryB
	expected["C"] = EntryC
	assertEquals(t, expected, sut)
	assertEquals(t, 2, len(BC))
}
