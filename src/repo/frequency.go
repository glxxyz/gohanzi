package repo

import (
	"encoding/csv"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

var FrequencyOrderedChars = []rune{}
var FrequencyOrderedWords = []string{}

func SortCharsHighestFrequencyFirst(chars []rune) {
	sort.Slice(chars, highestFreqCharComparator(chars))
}

func SortWordsHighestFrequencyFirst(words []string) {
	sort.Slice(words, highestFreqWordComparator(words))
}

func ParseWordFrequencyFile(dataDir string) {
	csvFile, err := os.Open(path.Join(dataDir, "SUBTLEX-CH-WF.txt"))
	if err != nil {
		panic(err)
	}
	if err := parseWordFrequency(csvFile); err != nil {
		panic(err)
	}
	if err := csvFile.Close(); err != nil {
		panic(err)
	}
}

func ParseCharFrequencyFile(dataDir string) {
	csvFile, err := os.Open(path.Join(dataDir, "SUBTLEX-CH-CHR.txt"))
	if err != nil {
		panic(err)
	}
	if err := parseCharFrequency(csvFile); err != nil {
		panic(err)
	}
	if err := csvFile.Close(); err != nil {
		panic(err)
	}
}

func highestFreqWordComparator(words []string) func(int, int) bool {
	return func(i, j int) bool {
		leftEntry := findEntry(words[i], "")
		rightEntry := findEntry(words[j], "")
		if leftEntry == nil {
			return false
		}
		if rightEntry == nil {
			return true
		}
		return leftEntry.WordFrequency > rightEntry.WordFrequency
	}
}

func highestFreqCharComparator(chars []rune) func(int, int) bool {
	return func(i, j int) bool {
		leftEntry := findEntry(string(chars[i]), "")
		rightEntry := findEntry(string(chars[j]), "")
		if leftEntry == nil {
			return false
		}
		if rightEntry == nil {
			return true
		}
		return leftEntry.WordFrequency > rightEntry.WordFrequency
	}
}

func parseWordFrequency(csvFile *os.File) error {
	reader := csv.NewReader(csvFile)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		return err
	}

	line := 0
	for _, fields := range csvData {
		// skip first 4 lines
		if line > 3 {
			simplified := fields[0]
			if strings.Contains(simplified, "\uFEFF") {
				simplified = strings.Replace(simplified, "\uFEFF", "", -1)
			}
			frequency, err := strconv.Atoi(fields[1])
			if err != nil {
				return err
			}

			if entries, found := HanziIndex[simplified]; found {
				for entry, _ := range entries {
					entry.WordFrequency = frequency
				}
			} else {
				entry := createEntry(simplified, "", true)
				entry.WordFrequency = frequency
			}
			FrequencyOrderedWords = append(FrequencyOrderedWords, simplified)
		}
		line += 1
	}

	SortWordsHighestFrequencyFirst(FrequencyOrderedWords)

	return nil
}

func parseCharFrequency(csvFile *os.File) error {
	reader := csv.NewReader(csvFile)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	line := 0
	for _, fields := range csvData {
		// skip first 4 lines
		if line > 3 {
			simplified := fields[0]
			if strings.Contains(simplified, "\uFEFF") {
				simplified = strings.Replace(simplified, "\uFEFF", "", -1)
			}
			frequency, err := strconv.Atoi(fields[1])
			if err != nil {
				panic(err)
			}
			runeCount := utf8.RuneCountInString(simplified)
			if runeCount != 1 {
				log.Panicf("Expected a single rune but found %d runes: %s", runeCount, simplified)
			}
			char := []rune(simplified)[0]

			if entries, found := CharIndex[char]; found {
				for entry, _ := range entries {
					entry.CharFrequency = frequency
				}
			} else {
				entry := createEntry(simplified, "", true)
				entry.WordFrequency = frequency
			}
			FrequencyOrderedChars = append(FrequencyOrderedChars, char)
		}
		line += 1
	}

	SortCharsHighestFrequencyFirst(FrequencyOrderedChars)

	return nil
}
