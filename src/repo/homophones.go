package repo

import "github.com/glxxyz/gohanzi/containers"

type Homophone struct {
	Pinyin  string
	Members []HomophoneMember
}

type HomophoneMember struct {
	Hanzi          string
	HskLevel       containers.HskLevel
	Definition     string
}

func BuildHomophones(numChars int8, matchTones bool, hskOnly bool) []Homophone {
	homophoneIndexMap := TonelessPinyinHomophones
	if matchTones {
		homophoneIndexMap = TonedPinyinHomophones
	}
	if homophoneIndex, found := homophoneIndexMap[numChars]; found {
		return homophonesFromMap(homophoneIndex, hskOnly)
	}
	return []Homophone{}
}

func homophonesFromMap(homophoneIndex containers.StringIndex, hskOnly bool) []Homophone {
	homophones := make([]Homophone, 0)
	for pinyin, entries := range homophoneIndex {
		homophones = append(homophones, buildHomophone(pinyin, entries, hskOnly))
	}
	return homophones
}

func buildHomophone(pinyin string, entries containers.EntrySet, hskOnly bool) Homophone {
	homophone := Homophone{
		Pinyin:  pinyin,
		Members: []HomophoneMember{},
	}
	for entry, _ := range entries {
		if !hskOnly || entry.GetHskWordLevel(Hsk2012) > 0 {
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
