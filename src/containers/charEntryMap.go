package containers

type CharEntryMap map[rune]*Entry

func (m CharEntryMap) Contains(value rune) bool {
	_, found := m[value]
	return found
}

func (m CharEntryMap) Difference(other CharEntryMap) CharEntryMap {
	difference := CharEntryMap{}
	for key, val := range m {
		if _, found := other[key]; !found {
			difference[key] = val
		}
	}
	return difference
}

func (m CharEntryMap) Intersection(other CharEntryMap) CharEntryMap {
	intersection := CharEntryMap{}
	for key, val := range m {
		if _, found := other[key]; found {
			intersection[key] = val
		}
	}
	return intersection
}

func (m CharEntryMap) SubsetOf(other CharEntryMap) bool {
	if len(m) > len(other) {
		return false
	}
	for key, _ := range m {
		if _, found := other[key]; !found {
			return false
		}
	}
	return true
}

func (m CharEntryMap) Union(other CharEntryMap) CharEntryMap {
	union := CharEntryMap{}
	for key, val := range m {
		union[key] = val
	}
	for key, val := range other {
		union[key] = val
	}
	return union
}

func (m CharEntryMap) AddAll(other CharEntryMap) {
	for key, val := range other {
		m[key] = val
	}
}
