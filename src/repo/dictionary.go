package repo

import (
	"bufio"
	"github.com/glxxyz/gohanzi/containers"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"unicode/utf8"
)

// A bunch of different ways of indexing the dictionary entries
var HanziIndex = containers.StringIndex{}
var CharIndex = containers.CharIndex{}
var RadicalIndex = containers.CharIndex{} // Index of character entries that use a given radical
var PinyinIndex = containers.StringIndex{}
var EnglishIndex = containers.StringIndex{}
var TonelessPinyinHomophones = map[int8]containers.StringIndex{}
var TonedPinyinHomophones = map[int8]containers.StringIndex{}
var ComponentIndex = map[rune]containers.CharSet{}
var ComposesIndex = map[rune]containers.CharSet{}

func addOrUpdateEntry(simplified string, traditional string, pinyin string, definition string, isWord bool) *containers.Entry {
	entry := findEntry(simplified, pinyin)
	if entry == nil {
		entry = createEntry(simplified, pinyin, isWord)
	}
	updateEntry(entry, traditional, definition, isWord)
	return entry
}

func findEntry(simplified string, pinyin string) *containers.Entry {
	if entries, found := HanziIndex[simplified]; found {
		for testEntry, _ := range entries {
			if pinyin == "" || pinyin == testEntry.Pinyin {
				return testEntry
			}
		}
	}
	return nil
}

func createEntry(simplified string, pinyin string, isWord bool) *containers.Entry {
	entry := &containers.Entry{
		Simplified: simplified,
		Pinyin:     pinyin,
		IsWord:     isWord,
	}
	HanziIndex.Add(simplified, entry)
	for _, char := range simplified {
		CharIndex.Add(char, entry)
	}
	tonelessSyllables, tonelessPinyin := parsePinyinNumToneless(pinyin)
	if isWord {
		createHomophoneEntry(pinyin, entry, tonelessSyllables, tonelessPinyin)
	}
	PinyinIndex.Add(pinyin, entry)
	PinyinIndex.Add(tonelessPinyin, entry)
	return entry
}

func createHomophoneEntry(pinyin string, entry *containers.Entry, tonelessSyllables []string, tonelessPinyin string) {
	syllableCount := int8(len(tonelessSyllables))
	if _, found := TonedPinyinHomophones[syllableCount]; !found {
		TonedPinyinHomophones[syllableCount] = containers.StringIndex{}
	}
	TonedPinyinHomophones[syllableCount].Add(pinyin, entry)
	if _, found := TonelessPinyinHomophones[syllableCount]; !found {
		TonelessPinyinHomophones[syllableCount] = containers.StringIndex{}
	}
	TonelessPinyinHomophones[syllableCount].Add(tonelessPinyin, entry)
}

func updateEntry(entry *containers.Entry, traditional string, definition string, isWord bool) {
	if entry.Traditional == "" {
		entry.Traditional = traditional
	}
	if definition != "" {
		if entry.Definition == "" {
			entry.Definition = definition
		} else {
			entry.Definition = entry.Definition + "\n" + definition
		}
		for _, word := range englishWords(definition) {
			EnglishIndex.Add(word, entry)
		}
		entry.ShortDefinition = entry.Definition
		if utf8.RuneCountInString(entry.Definition) > 80 {
			entry.ShortDefinition = string([]rune(entry.Definition)[:77]) + "..."
		}
	}
	if !entry.IsWord && isWord {
		tonelessSyllables, tonelessPinyin := parsePinyinNumToneless(entry.Pinyin)
		createHomophoneEntry(entry.Pinyin, entry, tonelessSyllables, tonelessPinyin)
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

var parseCcCeDictLineRegex = regexp.MustCompile(`^(\S+)\s+(\S+).*?\[(.*?)] /(.*)/$`)

func ParseCcCeDict(dataDir string) {
	fileName := path.Join(dataDir, "cedict_ts.u8")

	dictFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer dictFile.Close()

	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '#' {
			continue
		}
		matches := parseCcCeDictLineRegex.FindAllStringSubmatch(line, -1)

		traditional := matches[0][1]
		simplified := matches[0][2]
		pinyinNum := matches[0][3]
		definition := matches[0][4]
		definition = strings.Replace(definition, "/", "\n", -1)

		simplifiedChars := []rune(simplified)
		traditionalChars := []rune(traditional)
		pinyinVariants := strings.Split(pinyinNum, ",")
		for _, pinyin := range pinyinVariants {
			// don't parse the pinyin, just need to split it
			// syllables, cleanPinyin := parsePinyinNumTones(pinyin)
			syllables := strings.Split(pinyin, " ")
			// don't validate- fails at e.g. 21 = er4 shi2 yi1
			// validateEntry(simplified, traditional, syllables, pinyin)
			for index, char := range simplifiedChars {
				charPinyin := ""
				// Only use the per-character pinyin if we can reliably map the pinyin to characters, we can't
				// validate if we have e.g. 21 = er4 shi2 yi1
				if len(syllables) == len(traditionalChars) {
					charPinyin = syllables[index]
				}
				addOrUpdateEntry(string(char), string(traditionalChars[index]), charPinyin, "", false)
			}
			addOrUpdateEntry(simplified, traditional, pinyin, definition, true)
		}
	}
}

func validateEntry(simplified string, traditional string, syllables []string, pinyin string) {
	simplifiedCount := utf8.RuneCountInString(simplified)
	traditionalCount := utf8.RuneCountInString(traditional)
	syllablesCount := len(syllables)
	if simplifiedCount != traditionalCount || syllablesCount>0 && traditionalCount != syllablesCount {
		log.Panicf(
			"Character/pinyin counts do not agree for entry: Simplified(%v: [%v]), Traditional(%v: [%v]), or Pinyin(%v: [%v] from [%v])",
			simplifiedCount,
			simplified,
			traditionalCount,
			traditional,
			syllablesCount,
			strings.Join(syllables, ","),
			pinyin)
	}
}
