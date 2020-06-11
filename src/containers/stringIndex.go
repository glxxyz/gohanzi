package containers

type StringIndex map[string]EntrySet

func (index StringIndex) Add(word string, entry *Entry) {
	set, found := index[word]; if found {
		set[entry] = exists
	}
	set = EntrySet{}
	index[word] = set
	set[entry] = exists
}

func (index StringIndex) AddAll(other StringIndex) {
	for word, set := range other {
		for entry, _ := range set {
			index.Add(word, entry)
		}
	}
}
