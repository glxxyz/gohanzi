package containers

type CharSet map[rune]nothing

func (set CharSet) Add(value rune) {
	set[value] = exists
}

func (set CharSet) Contains(value rune) bool {
	_, found := set[value]
	return found
}

func (set CharSet) Difference(other CharSet) CharSet {
	difference := CharSet{}
	for key, val := range set {
		if _, found := other[key]; !found {
			difference[key] = val
		}
	}
	return difference
}

func (set CharSet) Intersection(other CharSet) CharSet {
	intersection := CharSet{}
	for key, val := range set {
		if _, found := other[key]; found {
			intersection[key] = val
		}
	}
	return intersection
}

func (set CharSet) SubsetOf(other CharSet) bool {
	if len(set) > len(other) {
		return false
	}
	for key, _ := range set {
		if _, found := other[key]; !found {
			return false
		}
	}
	return true
}

func (set CharSet) Union(other CharSet) CharSet {
	union := CharSet{}
	for key, _ := range set {
		union[key] = exists
	}
	for key, _ := range other {
		union[key] = exists
	}
	return union
}

func (set CharSet) AddAll(other CharSet) {
	for key, _ := range other {
		set[key] = exists
	}
}
