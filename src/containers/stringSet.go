package containers

type StringSet map[string]nothing

func (set StringSet) Add(value string) {
	set[value] = exists
}

func (set StringSet) Contains(value string) bool {
	_, found := set[value]
	return found
}

func (set StringSet) Difference(other StringSet) StringSet {
	difference := StringSet{}
	for key, val := range set {
		if _, found := other[key]; !found {
			difference[key] = val
		}
	}
	return difference
}

func (set StringSet) Intersection(other StringSet) StringSet {
	intersection := StringSet{}
	for key, val := range set {
		if _, found := other[key]; found {
			intersection[key] = val
		}
	}
	return intersection
}

func (set StringSet) SubsetOf(other StringSet) bool {
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

func (set StringSet) Union(other StringSet) StringSet {
	union := StringSet{}
	for key, _ := range set {
		union[key] = exists
	}
	for key, _ := range other {
		union[key] = exists
	}
	return union
}

func (set StringSet) AddAll(other StringSet) {
	for key, _ := range other {
		set[key] = exists
	}
}
