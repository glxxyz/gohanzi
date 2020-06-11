package containers

type StringEntryMap map[string]*Entry

func (m StringEntryMap) Contains(value string) bool {
	_, found := m[value]
	return found
}

func (m StringEntryMap) Difference(other StringEntryMap) StringEntryMap {
	difference := StringEntryMap{}
	for key, val := range m {
		if _, found := other[key]; !found {
			difference[key] = val
		}
	}
	return difference
}

func (m StringEntryMap) Intersection(other StringEntryMap) StringEntryMap {
	intersection := StringEntryMap{}
	for key, val := range m {
		if _, found := other[key]; found {
			intersection[key] = val
		}
	}
	return intersection
}

func (m StringEntryMap) SubsetOf(other StringEntryMap) bool {
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

func (m StringEntryMap) Union(other StringEntryMap) StringEntryMap {
	union := StringEntryMap{}
	for key, val := range m {
		union[key] = val
	}
	for key, val := range other {
		union[key] = val
	}
	return union
}

func (m StringEntryMap) AddAll(other StringEntryMap) {
	for key, val := range other {
		m[key] = val
	}
}
