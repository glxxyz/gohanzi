package containers_test

import (
	"github.com/glxxyz/gohanzi/containers"
	"testing"
)

func TestCharSet_AddContains(t *testing.T) {
	sut := containers.CharSetOf('A', 'B')
	assert(t, sut.Contains('A'), "contains A")
	assert(t, sut.Contains('B'), "contains B")
	assert(t, !sut.Contains('C'), "doesn't contain C")
	sut.Add('C')
	assert(t, sut.Contains('C'), "contains C")
}

func TestCharSet_Difference(t *testing.T) {
	AB := containers.CharSetOf('A', 'B')
	BC := containers.CharSetOf('B', 'C')
	result := AB.Difference(BC)
	equals(t, containers.CharSetOf('A'), result)
}

func TestCharSet_Intersection(t *testing.T) {
	AB := containers.CharSetOf('A', 'B')
	BC := containers.CharSetOf('B', 'C')
	result := AB.Intersection(BC)
	equals(t, containers.CharSetOf('B'), result)
}

func TestCharSet_SubSetOf(t *testing.T) {
	ABC := containers.CharSetOf('A', 'B', 'C')
	BC := containers.CharSetOf('B', 'C')
	ABZ := containers.CharSetOf('A', 'B', 'Z')
	assert(t, ABC.SubsetOf(ABC), "ABC is a subset of ABC")
	assert(t, BC.SubsetOf(ABC), "BC is a subset of ABC")
	assert(t, !ABC.SubsetOf(BC), "ABC is not a subset of BC")
	assert(t, !ABC.SubsetOf(ABZ), "ABC is not a subset of ABZ")
}

func TestCharSet_Union(t *testing.T) {
	AB := containers.CharSetOf('A', 'B')
	BC := containers.CharSetOf('B', 'C')
	result := AB.Union(BC)
	equals(t, containers.CharSetOf('A', 'B'), AB)
	equals(t, containers.CharSetOf('B', 'C'), BC)
	equals(t, containers.CharSetOf('A', 'B', 'C'), result)
}

func TestCharSet_AddAll(t *testing.T) {
	sut := containers.CharSetOf('A', 'B')
	BC := containers.CharSetOf('B', 'C')
	sut.AddAll(BC)
	equals(t, containers.CharSetOf('A', 'B', 'C'), sut)
	equals(t, containers.CharSetOf('B', 'C'), BC)
}
