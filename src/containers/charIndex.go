package containers

type CharIndex map[rune]EntrySet

func (index CharIndex) Add(char rune, entry *Entry) {
	if set, found := index[char]; found {
		set[entry] = exists
		return
	}
	index[char] = EntrySet{entry: exists}
}

func (index CharIndex) AddAll(other CharIndex) {
	for char, set := range other {
		thisSet, found := index[char]
		if !found {
			thisSet = EntrySet{}
			index[char] = thisSet
		}
		for entry, _ := range set {
			thisSet[entry] = exists
		}
	}
}
