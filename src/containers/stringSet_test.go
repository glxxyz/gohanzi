package containers_test

import (
	"github.com/glxxyz/gohanzi/containers"
	"testing"
)

func TestStringSet_AddContains(t *testing.T) {
	sut := containers.StringSetOf("A", "B")
	assert(t, sut.Contains("A"), "contains A")
	assert(t, sut.Contains("B"), "contains B")
	assert(t, !sut.Contains("C"), "doesn't contain C")
	sut.Add("C")
	assert(t, sut.Contains("C"), "contains C")
}

func TestStringSet_Difference(t *testing.T) {
	AB := containers.StringSetOf("A", "B")
	BC := containers.StringSetOf("B", "C")
	result := AB.Difference(BC)
	equals(t, containers.StringSetOf("A"), result)
}

func TestStringSet_Intersection(t *testing.T) {
	AB := containers.StringSetOf("A", "B")
	BC := containers.StringSetOf("B", "C")
	result := AB.Intersection(BC)
	equals(t, containers.StringSetOf("B"), result)
}

func TestStringSet_SubSetOf(t *testing.T) {
	ABC := containers.StringSetOf("A", "B", "C")
	BC := containers.StringSetOf("B", "C")
	ABZ := containers.StringSetOf("A", "B", "Z")
	assert(t, ABC.SubsetOf(ABC), "ABC is a subset of ABC")
	assert(t, BC.SubsetOf(ABC), "BC is a subset of ABC")
	assert(t, !ABC.SubsetOf(BC), "ABC is not a subset of BC")
	assert(t, !ABC.SubsetOf(ABZ), "ABC is not a subset of ABZ")
}

func TestStringSet_Union(t *testing.T) {
	AB := containers.StringSetOf("A", "B")
	BC := containers.StringSetOf("B", "C")
	result := AB.Union(BC)
	equals(t, containers.StringSetOf("A", "B"), AB)
	equals(t, containers.StringSetOf("B", "C"), BC)
	equals(t, containers.StringSetOf("A", "B", "C"), result)
}

func TestStringSet_AddAll(t *testing.T) {
	sut := containers.StringSetOf("A", "B")
	BC := containers.StringSetOf("B", "C")
	sut.AddAll(BC)
	equals(t, containers.StringSetOf("A", "B", "C"), sut)
	equals(t, containers.StringSetOf("B", "C"), BC)
}
