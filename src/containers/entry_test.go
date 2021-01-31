package containers_test

import (
	"github.com/glxxyz/gohanzi/containers"
	"testing"
)

const HskVersionA containers.HskVersion = 1
const HskVersionB containers.HskVersion = 2
const NoHskLevel containers.HskLevel = 0
const HskLevel1 containers.HskLevel = 1
const HskLevel2 containers.HskLevel = 2
const HskLevel3 containers.HskLevel = 3

func TestEntry_SetGreaterWordLevel(t *testing.T) {
	sut := containers.Entry{}
	equals(t, NoHskLevel, sut.GetHskWordLevel(HskVersionA))
	assert(t, !sut.IsHskWordLevel(HskVersionA, HskLevel3), "Version A of this word is not level 3")
	sut.SetHskWordLevel(HskVersionA, HskLevel2)
	sut.SetHskWordLevel(HskVersionA, HskLevel3)
	equals(t, HskLevel2, sut.GetHskWordLevel(HskVersionA))
	assert(t, sut.IsHskWordLevel(HskVersionA, HskLevel2), "Version A of this word is level 2")
	assert(t, !sut.IsHskWordLevel(HskVersionA, HskLevel3), "Version A of this word is not level 3")
	equals(t, NoHskLevel, sut.GetHskWordLevel(HskVersionB))
	assert(t, !sut.IsHskWordLevel(HskVersionB, HskLevel2), "Version A of this word is not level 2")
}

func TestEntry_SetLesserWordLevel(t *testing.T) {
	sut := containers.Entry{}
	equals(t, NoHskLevel, sut.GetHskWordLevel(HskVersionA))
	assert(t, !sut.IsHskWordLevel(HskVersionA, HskLevel3), "Version A of this word is not level 3")
	sut.SetHskWordLevel(HskVersionA, HskLevel3)
	sut.SetHskWordLevel(HskVersionA, HskLevel2)
	equals(t, HskLevel2, sut.GetHskWordLevel(HskVersionA))
	assert(t, sut.IsHskWordLevel(HskVersionA, HskLevel2), "Version A of this word is Level 2")
	assert(t, !sut.IsHskWordLevel(HskVersionA, HskLevel3), "Version A of this word is not level 3")
	equals(t, NoHskLevel, sut.GetHskWordLevel(HskVersionB))
	assert(t, !sut.IsHskWordLevel(HskVersionB, HskLevel2), "Version A of this word is not level 2")
}

func TestEntry_SetTwoWordLevels(t *testing.T) {
	sut := containers.Entry{}
	equals(t, NoHskLevel, sut.GetHskWordLevel(HskVersionA))
	equals(t, NoHskLevel, sut.GetHskWordLevel(HskVersionB))
	assert(t, !sut.IsHskWordLevel(HskVersionA, HskLevel1), "Version A of this word is not Level 1")
	assert(t, !sut.IsHskWordLevel(HskVersionB, HskLevel2), "Version B of this word is not Level 2")
	sut.SetHskWordLevel(HskVersionA, HskLevel1)
	sut.SetHskWordLevel(HskVersionB, HskLevel2)
	equals(t, HskLevel1, sut.GetHskWordLevel(HskVersionA))
	equals(t, HskLevel2, sut.GetHskWordLevel(HskVersionB))
	assert(t, sut.IsHskWordLevel(HskVersionA, HskLevel1), "Version A of this word is Level 1")
	assert(t, sut.IsHskWordLevel(HskVersionB, HskLevel2), "Version A of this word is Level 2")
	assert(t, !sut.IsHskWordLevel(HskVersionA, HskLevel2), "Version A of this word is not level 2")
}

func TestEntry_SetGreaterCharLevel(t *testing.T) {
	sut := containers.Entry{}
	equals(t, NoHskLevel, sut.GetHskCharLevel(HskVersionA))
	assert(t, !sut.IsHskCharLevel(HskVersionA, HskLevel3), "Version A of this char is not level 3")
	sut.SetHskCharLevel(HskVersionA, HskLevel2)
	sut.SetHskCharLevel(HskVersionA, HskLevel3)
	equals(t, HskLevel2, sut.GetHskCharLevel(HskVersionA))
	assert(t, sut.IsHskCharLevel(HskVersionA, HskLevel2), "Version A of this char is level 2")
	assert(t, !sut.IsHskCharLevel(HskVersionA, HskLevel3), "Version A of this char is not level 3")
	equals(t, NoHskLevel, sut.GetHskCharLevel(HskVersionB))
	assert(t, !sut.IsHskCharLevel(HskVersionB, HskLevel2), "Version A of this char is not level 2")
}

func TestEntry_SetLesserCharLevel(t *testing.T) {
	sut := containers.Entry{}
	equals(t, NoHskLevel, sut.GetHskCharLevel(HskVersionA))
	assert(t, !sut.IsHskCharLevel(HskVersionA, HskLevel3), "Version A of this char is not level 3")
	sut.SetHskCharLevel(HskVersionA, HskLevel3)
	sut.SetHskCharLevel(HskVersionA, HskLevel2)
	equals(t, HskLevel2, sut.GetHskCharLevel(HskVersionA))
	assert(t, sut.IsHskCharLevel(HskVersionA, HskLevel2), "Version A of this char is Level 2")
	assert(t, !sut.IsHskCharLevel(HskVersionA, HskLevel3), "Version A of this char is not level 3")
	equals(t, NoHskLevel, sut.GetHskCharLevel(HskVersionB))
	assert(t, !sut.IsHskCharLevel(HskVersionB, HskLevel2), "Version A of this char is not level 2")
}

func TestEntry_SetTwoCharLevels(t *testing.T) {
	sut := containers.Entry{}
	equals(t, NoHskLevel, sut.GetHskCharLevel(HskVersionA))
	equals(t, NoHskLevel, sut.GetHskCharLevel(HskVersionB))
	assert(t, !sut.IsHskCharLevel(HskVersionA, HskLevel1), "Version A of this char is not Level 1")
	assert(t, !sut.IsHskCharLevel(HskVersionB, HskLevel2), "Version B of this char is not Level 2")
	sut.SetHskCharLevel(HskVersionA, HskLevel1)
	sut.SetHskCharLevel(HskVersionB, HskLevel2)
	equals(t, HskLevel1, sut.GetHskCharLevel(HskVersionA))
	equals(t, HskLevel2, sut.GetHskCharLevel(HskVersionB))
	assert(t, sut.IsHskCharLevel(HskVersionA, HskLevel1), "Version A of this char is Level 1")
	assert(t, sut.IsHskCharLevel(HskVersionB, HskLevel2), "Version A of this char is Level 2")
	assert(t, !sut.IsHskCharLevel(HskVersionA, HskLevel2), "Version A of this char is not level 2")
}