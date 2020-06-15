package containers

type Entry struct {
	Simplified      string
	Traditional     string
	Pinyin          string // tone numbers, spaces, v for u-umlaut e.g. "lv3 hao3 ba"
	Definition      string
	ShortDefinition string
	Radical         rune
	IsWord          bool
	hskWordLevel    map[HskVersion]HskLevel
	hskCharLevel    map[HskVersion]HskLevel
}

type HskVersion int8
type HskLevel int8

func (entry *Entry) SetHskWordLevel(version HskVersion, level HskLevel) {
	if entry.hskWordLevel == nil {
		entry.hskWordLevel = map[HskVersion]HskLevel{}
	}
	oldLevel, found := entry.hskWordLevel[version]
	if found && oldLevel > 0 && oldLevel <= level {
		return
	}
	entry.hskWordLevel[version] = level
}

func (entry *Entry) GetHskWordLevel(version HskVersion) HskLevel {
	if entry.hskWordLevel == nil {
		return 0
	}
	if level, found := entry.hskWordLevel[version]; found {
		return level
	}
	return 0
}

func (entry *Entry) IsHskWordLevel(version HskVersion, level HskLevel) bool {
	if entry.hskWordLevel == nil {
		return false
	}
	if actualLevel, found := entry.hskWordLevel[version]; found {
		return level == actualLevel
	}
	return false
}

func (entry *Entry) SetHskCharLevel(version HskVersion, level HskLevel) {
	if entry.hskCharLevel == nil {
		entry.hskCharLevel = map[HskVersion]HskLevel{}
	}
	if oldLevel, found := entry.hskCharLevel[version]; found && oldLevel > 0 && oldLevel <= level {
		return
	}
	entry.hskCharLevel[version] = level
}

func (entry *Entry) GetHskCharLevel(version HskVersion) HskLevel {
	if entry.hskCharLevel == nil {
		return 0
	}
	if level, found := entry.hskCharLevel[version]; found {
		return level
	}
	return 0
}

func (entry *Entry) IsHskCharLevel(version HskVersion, level HskLevel) bool {
	if entry.hskCharLevel == nil {
		return false
	}
	if actualLevel, found := entry.hskCharLevel[version]; found {
		return level == actualLevel
	}
	return false
}
