package containers_test

import (
	"github.com/glxxyz/gohanzi/containers"
	"testing"
)

var EntryA = &containers.Entry{Simplified: "A"}
var EntryB = &containers.Entry{Simplified: "B"}
var EntryC = &containers.Entry{Simplified: "C"}
var EntryZ = &containers.Entry{Simplified: "Z"}

func TestEntrySet_AddContains(t *testing.T) {
	sut := containers.EntrySetOf(EntryA, EntryB)
	assert(t, sut.Contains(EntryA), "contains A")
	assert(t, sut.Contains(EntryB), "contains B")
	assert(t, !sut.Contains(EntryC), "doesn't contain C")
	sut.Add(EntryC)
	assert(t, sut.Contains(EntryC), "contains C")
}

func TestEntrySet_Difference(t *testing.T) {
	AB := containers.EntrySetOf(EntryA, EntryB)
	BC := containers.EntrySetOf(EntryB, EntryC)
	result := AB.Difference(BC)
	assertEquals(t, containers.EntrySetOf(EntryA), result)
}

func TestEntrySet_Intersection(t *testing.T) {
	AB := containers.EntrySetOf(EntryA, EntryB)
	BC := containers.EntrySetOf(EntryB, EntryC)
	result := AB.Intersection(BC)
	assertEquals(t, containers.EntrySetOf(EntryB), result)
}

func TestEntrySet_SubSetOf(t *testing.T) {
	ABC := containers.EntrySetOf(EntryA, EntryB, EntryC)
	BC := containers.EntrySetOf(EntryB, EntryC)
	ABZ := containers.EntrySetOf(EntryA, EntryB, EntryZ)
	assert(t, ABC.SubsetOf(ABC), "ABC is a subset of ABC")
	assert(t, BC.SubsetOf(ABC), "BC is a subset of ABC")
	assert(t, !ABC.SubsetOf(BC), "ABC is not a subset of BC")
	assert(t, !ABC.SubsetOf(ABZ), "ABC is not a subset of ABZ")
}

func TestEntrySet_Union(t *testing.T) {
	AB := containers.EntrySetOf(EntryA, EntryB)
	BC := containers.EntrySetOf(EntryB, EntryC)
	result := AB.Union(BC)
	assertEquals(t, containers.EntrySetOf(EntryA, EntryB), AB)
	assertEquals(t, containers.EntrySetOf(EntryB, EntryC), BC)
	assertEquals(t, containers.EntrySetOf(EntryA, EntryB, EntryC), result)
}

func TestEntrySet_AddAll(t *testing.T) {
	sut := containers.EntrySetOf(EntryA, EntryB)
	BC := containers.EntrySetOf(EntryB, EntryC)
	sut.AddAll(BC)
	assertEquals(t, containers.EntrySetOf(EntryA, EntryB, EntryC), sut)
	assertEquals(t, containers.EntrySetOf(EntryB, EntryC), BC)
}
