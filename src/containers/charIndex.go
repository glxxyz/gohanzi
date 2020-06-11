package containers

type CharIndex map[rune]EntrySet

func (index CharIndex) Add(char rune, entry *Entry) {
	set, found := index[char]; if found {
		set[entry] = exists
	}
	set = EntrySet{}
	index[char] = set
	set[entry] = exists
}

func (index CharIndex) AddAll(other CharIndex) {
	for char, set := range other {
		for entry, _ := range set {
			index.Add(char, entry)
		}
	}
}
