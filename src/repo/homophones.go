package repo

import "github.com/glxxyz/gohanzi/containers"

type Homophone struct {
	Pinyin  string
	Members []HomophoneMember
}

type HomophoneMember struct {
	Hanzi      string
	HskLevel   containers.HskLevel
	Definition string
}

func BuildHomophones(numChars int8, matchTones bool, hskVersion containers.HskVersion) []Homophone {
	homophoneIndexMap := TonelessPinyinHomophones
	if matchTones {
		homophoneIndexMap = TonedPinyinHomophones
	}
	if homophoneIndex, found := homophoneIndexMap[numChars]; found {
		return homophonesFromMap(homophoneIndex, hskVersion)
	}
	return []Homophone{}
}

func homophonesFromMap(homophoneIndex containers.StringIndex, hskVersion containers.HskVersion) []Homophone {
	homophones := make([]Homophone, 0)
	for pinyin, entries := range homophoneIndex {
		homophone := buildHomophone(pinyin, entries, hskVersion)
		if len(homophone.Members) > 1 {
			homophones = append(homophones, homophone)
		}
	}
	return homophones
}

func buildHomophone(pinyin string, entries containers.EntrySet, hskVersion containers.HskVersion) Homophone {
	homophone := Homophone{
		Pinyin:  pinyin,
		Members: []HomophoneMember{},
	}
	for entry, _ := range entries {
		if hskVersion == HskNone || entry.GetHskWordLevel(hskVersion) > 0 {
			homophone.Members = append(homophone.Members, buildHomophoneMember(entry))
		}
	}
	return homophone
}

func buildHomophoneMember(entry *containers.Entry) HomophoneMember {
	member := HomophoneMember{
		Hanzi:      entry.Simplified,
		HskLevel:   entry.GetHskWordLevel(Hsk2012),
		Definition: entry.ShortDefinition,
	}
	return member
}
