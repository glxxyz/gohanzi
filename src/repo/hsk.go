package repo

import (
	"encoding/csv"
	"github.com/glxxyz/gohanzi/containers"
	"os"
	"path"
	"strconv"
	"strings"
)

var HskWords = map[containers.HskVersion]map[containers.HskLevel]containers.StringEntryMap{}
var HskChars = map[containers.HskVersion]map[containers.HskLevel]containers.CharEntryMap{}

const Hsk1992 containers.HskVersion = 92
const Hsk2010 containers.HskVersion = 10
const Hsk2012 containers.HskVersion = 12
const Hsk2020 containers.HskVersion = 20

var HskVersionToLevels  = map[containers.HskVersion]containers.HskLevel{
	Hsk1992: 4,
	Hsk2010: 6,
	Hsk2012: 6,
	Hsk2020: 9,
}

func parseHskFiles(dataDir string) {
	hsk2012Words := containers.StringEntryMap{}
	hsk2012Chars := containers.CharEntryMap{}
	parseHsk2012File(hsk2012Words, hsk2012Chars, path.Join(dataDir, "HSK Official With Definitions 2012 L1.txt"), 1)
	parseHsk2012File(hsk2012Words, hsk2012Chars, path.Join(dataDir, "HSK Official With Definitions 2012 L2.txt"), 2)
	parseHsk2012File(hsk2012Words, hsk2012Chars, path.Join(dataDir, "HSK Official With Definitions 2012 L3.txt"), 3)
	parseHsk2012File(hsk2012Words, hsk2012Chars, path.Join(dataDir, "HSK Official With Definitions 2012 L4.txt"), 4)
	parseHsk2012File(hsk2012Words, hsk2012Chars, path.Join(dataDir, "HSK Official With Definitions 2012 L5.txt"), 5)
	parseHsk2012File(hsk2012Words, hsk2012Chars, path.Join(dataDir, "HSK Official With Definitions 2012 L6.txt"), 6)
	buildHskLevelToLists(hsk2012Words, hsk2012Chars, Hsk2012)

	hsk2010Words, hsk2010Chars := parseHsk2010File(path.Join(dataDir, "New_HSK_2010.csv"))
	buildHskLevelToLists(hsk2010Words, hsk2010Chars, Hsk2010)

	hsk1992Words, hsk1992Chars := parseHsk1992File(path.Join(dataDir, "oldhsk.csv"))
	buildHskLevelToLists(hsk1992Words, hsk1992Chars, Hsk1992)
}

func parseHsk2012File(hskWords containers.StringEntryMap, hskChars containers.CharEntryMap, fileName string, hskLevel containers.HskLevel) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, each := range csvData {
		simplified := each[0]
		if strings.Contains(simplified, "\uFEFF") {
			simplified = strings.Replace(simplified, "\uFEFF", "", -1)
		}
		traditional := each[1]
		pinyinNum := each[2]
		// pinyinTones := each[3]
		definition := each[4]
		processHskEntry(hskWords, hskChars, Hsk2012, hskLevel, simplified, traditional, pinyinNum, definition)
	}
}

func parseHsk2010File(fileName string) (containers.StringEntryMap, containers.CharEntryMap) {
	hskWords := containers.StringEntryMap{}
	hskChars := containers.CharEntryMap{}

	csvFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, each := range csvData {
		levelStr := each[0]
		if strings.Contains(levelStr, "\uFEFF") {
			levelStr = strings.Replace(levelStr, "\uFEFF", "", -1)
		}
		levelInt, err := strconv.Atoi(levelStr)
		if err != nil {
			panic(err)
		}
		hskLevel := containers.HskLevel(levelInt)
		simplified := each[1]
		pinyinNum := each[2]
		processHskEntry(hskWords, hskChars, Hsk2010, hskLevel, simplified, "", pinyinNum, "")
	}

	return hskWords, hskChars
}

func processHskEntry(
  hskWords containers.StringEntryMap,
  hskChars containers.CharEntryMap,
  hskVersion containers.HskVersion,
  hskLevel containers.HskLevel,
  simplified string,
  traditional string,
  pinyinNum string,
  definition string) {

	simplifiedChars := []rune(simplified)
	traditionalChars := []rune(traditional)
	traditionalValidate := traditional
	if traditionalValidate == "" {
		traditionalValidate = simplified
	}
	pinyinVariants := strings.Split(pinyinNum, ",")
	for _, pinyin := range pinyinVariants {
		syllables, cleanPinyin := parsePinyinNumTones(pinyin)
		validateEntry(simplified, traditionalValidate, syllables, pinyin)
		for index, char := range simplifiedChars {
			var entry *containers.Entry
			if traditional == "" {
				entry = addOrUpdateEntry(string(char), "", syllables[index], "", false)
			} else {
				entry = addOrUpdateEntry(string(char), string(traditionalChars[index]), syllables[index], "", false)
			}
			entry.SetHskCharLevel(hskVersion, hskLevel)
			hskChars[char] = entry
		}
		entry := addOrUpdateEntry(simplified, traditional, cleanPinyin, definition, true)
		entry.SetHskWordLevel(hskVersion, hskLevel)
		hskWords[simplified] = entry
	}
}

func parseHsk1992File(fileName string) (containers.StringEntryMap, containers.CharEntryMap) {
	hskWords := containers.StringEntryMap{}
	hskChars := containers.CharEntryMap{}

	csvFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, each := range csvData {
		levelInt, err := strconv.Atoi(each[3])
		if err != nil {
			panic(err)
		}
		hskLevel := containers.HskLevel(levelInt)
		simplified := each[0]
		if strings.Contains(simplified, "\uFEFF") {
			simplified = strings.Replace(simplified, "\uFEFF", "", -1)
		}
		if strings.Contains(simplified, "_") {
			simplified = strings.Replace(simplified, "_", "", -1)
		}
		pinyinNum := each[2]
		processHskEntry(hskWords, hskChars, Hsk1992, hskLevel, simplified, "", pinyinNum, "")
	}

	return hskWords, hskChars
}

func buildHskLevelToLists(hskWords containers.StringEntryMap, hskChars containers.CharEntryMap,	hskVersion containers.HskVersion) {
	levelToWords := map[containers.HskLevel]containers.StringEntryMap{}
	levelToChars := map[containers.HskLevel]containers.CharEntryMap{}

	for i :=containers.HskLevel(1); i<=HskVersionToLevels[hskVersion]; i++ {
		levelToWords[i] = containers.StringEntryMap{}
		levelToChars[i] = containers.CharEntryMap{}
	}

	for word, entry := range hskWords {
		hskWordLevel := entry.GetHskWordLevel(hskVersion)
		levelToWords[hskWordLevel][word] = entry
	}

	for char, entry := range hskChars {
		hskCharLevel := entry.GetHskCharLevel(hskVersion)
		levelToChars[hskCharLevel][char] = entry
	}

	// build sets of character/word ranges; e.g. words[13] is the
	// union of the words for HSK levels 1, 2, and 3.

	// I don't think this is still needed- leave it out for now
	/*
	for i:=containers.HskLevel(1); i<=HskVersionToLevels[hskVersion]-1; i++ {
		for j := i + 1; j <= HskVersionToLevels[hskVersion]; j++ {
			index := i * 10 + j
			levelToWords[index] = containers.StringEntryMap{}
			levelToChars[index] = containers.CharEntryMap{}
			for k := i; k <= j; k++ {
				levelToWords[index].AddAll(levelToWords[k])
				levelToChars[index].AddAll(levelToChars[k])
			}
		}
	}
	 */

	HskWords[hskVersion] = levelToWords
	HskChars[hskVersion] = levelToChars
}
