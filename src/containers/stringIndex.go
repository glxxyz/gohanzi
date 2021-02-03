package containers

type StringIndex map[string]EntrySet

func (index StringIndex) Add(word string, entry *Entry) {
	if set, found := index[word]; found {
		set.Add(entry)
		return
	}
	index[word] = EntrySet{entry: exists}
}

func (index StringIndex) AddAll(other StringIndex) {
	for word, set := range other {
		thisSet, found := index[word]
		if !found {
			thisSet = EntrySet{}
			index[word] = thisSet
		}
		for entry, _ := range set {
			thisSet.Add(entry)
		}
	}
}
