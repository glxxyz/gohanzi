package repo

import (
	"github.com/glxxyz/gohanzi/containers"
	"regexp"
	"strings"
)

// A bunch of different ways of indexing the dictionary entries
var HanziIndex = containers.StringIndex{}
var CharIndex = containers.CharIndex{}
var RadicalIndex = containers.CharIndex{}
var PinyinIndex = containers.StringIndex{}
var EnglishIndex = containers.StringIndex{}
var TonelessPinyinHomophones = map[int8]containers.StringIndex{}
var TonedPinyinHomophones = map[int8]containers.StringIndex{}
var ComponentIndex = map[rune]containers.CharSet{}
var ComposesIndex = map[rune]containers.CharSet{}

func addOrUpdateEntry(simplified string, traditional string, pinyin string, definition string) *containers.Entry {
	entry := findEntry(simplified, pinyin)
	if entry == nil {
		entry = createEntry(entry, simplified, pinyin)
	}
	updateEntry(entry, traditional, definition)
	return entry
}

func findEntry(simplified string, pinyin string,) *containers.Entry {
	entries, found := HanziIndex[simplified]
	if found {
		for testEntry, _ := range entries {
			if pinyin == testEntry.Pinyin {
				return testEntry
			}
		}
	}
	return nil
}

func createEntry(entry *containers.Entry, simplified string, pinyin string) *containers.Entry {
	entry = &containers.Entry{}
	entry.Simplified = simplified
	entry.Pinyin = pinyin
	HanziIndex.Add(simplified, entry)
	for _, char := range simplified {
		CharIndex.Add(char, entry)
	}
	tonelessSyllables := pinyinTonelessSyllables(pinyin)
	syllableCount := int8(len(tonelessSyllables))
	tonelessPinyin := strings.Join(tonelessSyllables, " ")
	_, found := TonedPinyinHomophones[syllableCount]; if !found {
		TonedPinyinHomophones[syllableCount] = containers.StringIndex{}
	}
	TonedPinyinHomophones[syllableCount].Add(pinyin, entry)
	_, found = TonelessPinyinHomophones[syllableCount]; if !found {
		TonelessPinyinHomophones[syllableCount] = containers.StringIndex{}
	}
	TonelessPinyinHomophones[syllableCount].Add(tonelessPinyin, entry)
	PinyinIndex.Add(pinyin, entry)
	PinyinIndex.Add(tonelessPinyin, entry)
	return entry
}

func updateEntry(entry *containers.Entry, traditional string, definition string) {
	if entry.Traditional == "" {
		entry.Traditional = traditional
	}
	if definition != "" {
		if entry.Definition == nil {
			entry.Definition = []string{}
		}
		entry.Definition = append(entry.Definition, definition)
		for _, word := range englishWords(definition) {
			EnglishIndex.Add(word, entry)
		}
		shortDefinition := strings.Join(entry.Definition, "\n")
		if len(shortDefinition) > 80 {
			shortDefinition = string([]rune(shortDefinition)[:77]) + "..."
		}
		entry.ShortDefinition = shortDefinition
	}
}

var englishWordRegEx = regexp.MustCompile(`([A-Za-z]+)`)

func englishWords(english string) (words []string) {
	result := make([]string, 0)
	var matches = englishWordRegEx.FindAllStringSubmatch(english, -1)
	for _, match := range matches {
		result = append(result, match[1])
	}
	return result
}
