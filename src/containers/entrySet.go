package containers

type EntrySet map[*Entry]nothing

func EntrySetOf(values ...*Entry) EntrySet {
	setOf := EntrySet{}
	for _, v := range values {
		setOf.Add(v)
	}
	return setOf
}

func (set EntrySet) Add(value *Entry) {
	set[value] = exists
}

func (set EntrySet) Contains(value *Entry) bool {
	_, found := set[value]
	return found
}

func (set EntrySet) Difference(other EntrySet) EntrySet {
	difference := EntrySet{}
	for key, val := range set {
		if _, found := other[key]; !found {
			difference[key] = val
		}
	}
	return difference
}

func (set EntrySet) Intersection(other EntrySet) EntrySet {
	intersection := EntrySet{}
	for key, val := range set {
		if _, found := other[key]; found {
			intersection[key] = val
		}
	}
	return intersection
}

func (set EntrySet) SubsetOf(other EntrySet) bool {
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

func (set EntrySet) Union(other EntrySet) EntrySet {
	union := EntrySet{}
	for key, _ := range set {
		union[key] = exists
	}
	for key, _ := range other {
		union[key] = exists
	}
	return union
}

func (set EntrySet) AddAll(other EntrySet) {
	for key, _ := range other {
		set[key] = exists
	}
}
