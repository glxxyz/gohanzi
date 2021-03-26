package repo

import "github.com/glxxyz/gohanzi/containers"

func ParseCharacterCompositionFile(dataDir string) {
	// ...
	initializeRadicals()
}

func initializeRadicals() {
	for _, entries := range CharIndex {
		for entry := range entries {
			if entry.Radical == 0 {
				continue
			}
			radicalEntry := findEntry(string(entry.Radical), "")
			if radicalEntry == nil {
				radicalEntry = createEntry(string(entry.Radical), "", false)
			}
			radicalUses, found := RadicalIndex[entry.Radical];
			if !found {
				radicalUses = make(containers.EntrySet)
				RadicalIndex[entry.Radical] = radicalUses
			}
			radicalUses.Add(entry)
		}
	}
}
