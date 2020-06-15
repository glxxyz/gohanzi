package repo

import (
	"encoding/csv"
	"github.com/glxxyz/gohanzi/containers"
	"log"
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

var HskVersionToLevels = map[containers.HskVersion]containers.HskLevel{
	Hsk1992: 4,
	Hsk2010: 6,
	Hsk2012: 6,
	Hsk2020: 9,
}

func parseHskFiles(dataDir string) {
	parseHskVersionFiles(dataDir, Hsk1992, ',', extractHsk1992, "oldhsk.csv")
	parseHskVersionFiles(dataDir, Hsk2010, ',', extractHsk2010, "New_HSK_2010.csv")
	processHsk2012(dataDir)
	parseHskVersionFiles(dataDir, Hsk2020, ',', extractHsk2020, "2020_vocab_list.csv")
}

func processHsk2012(dataDir string) {
	hskWords, hskChars := containers.StringEntryMap{}, containers.CharEntryMap{}
	parseHskFile(hskWords, hskChars, path.Join(dataDir, "HSK Official With Definitions 2012 L1.txt"), Hsk2012, '\t', genExtractHsk2012(1))
	parseHskFile(hskWords, hskChars, path.Join(dataDir, "HSK Official With Definitions 2012 L2.txt"), Hsk2012, '\t', genExtractHsk2012(2))
	parseHskFile(hskWords, hskChars, path.Join(dataDir, "HSK Official With Definitions 2012 L3.txt"), Hsk2012, '\t', genExtractHsk2012(3))
	parseHskFile(hskWords, hskChars, path.Join(dataDir, "HSK Official With Definitions 2012 L4.txt"), Hsk2012, '\t', genExtractHsk2012(4))
	parseHskFile(hskWords, hskChars, path.Join(dataDir, "HSK Official With Definitions 2012 L5.txt"), Hsk2012, '\t', genExtractHsk2012(5))
	parseHskFile(hskWords, hskChars, path.Join(dataDir, "HSK Official With Definitions 2012 L6.txt"), Hsk2012, '\t', genExtractHsk2012(6))
	buildHskLevelToLists(hskWords, hskChars, Hsk2012)
}

func extractHsk1992(fields []string) (containers.HskLevel, string, string, string, string) {
	levelInt, err := strconv.Atoi(fields[3])
	if err != nil {
		panic(err)
	}
	hskLevel := containers.HskLevel(levelInt)
	simplified := fields[0]
	if strings.Contains(simplified, "_") {
		simplified = strings.Replace(simplified, "_", "", -1)
	}
	pinyinNum := fields[2]
	return hskLevel, simplified, "", pinyinNum, ""
}

func extractHsk2010(fields []string) (containers.HskLevel, string, string, string, string) {
	levelStr := fields[0]
	levelInt, err := strconv.Atoi(levelStr)
	if err != nil {
		panic(err)
	}
	hskLevel := containers.HskLevel(levelInt)
	simplified := fields[1]
	pinyinNum := fields[2]
	return hskLevel, simplified, "", pinyinNum, ""
}

func genExtractHsk2012(hskLevel containers.HskLevel) func (fields []string) (
  hskLevel containers.HskLevel, simplified string, traditional string, pinyinNum string, definition string) {
	return func(fields []string) (containers.HskLevel, string, string, string, string) {
		simplified := fields[0]
		traditional := fields[1]
		pinyinNum := fields[2]
		// pinyinTones := each[3]
		definition := fields[4]
		return hskLevel, simplified, traditional, pinyinNum, definition
	}
}

func extractHsk2020(fields []string) (containers.HskLevel, string, string, string, string) {
	variants := strings.Split(fields[0], "丨")
	simplified := variants[0] // a few fields had shorter variants after the pipe
	var hskLevel containers.HskLevel
	switch fields[1] {
	case "elementary":
		hskLevel = 1
	case "intermediate":
		hskLevel = 2
	case "advanced":
		hskLevel = 3
	case "extra":
		hskLevel = 4
	default:
		log.Panicf("Unknown HSK 2020 level: %v", fields[1])
	}
	return hskLevel, simplified, "", "", ""
}

func parseHskVersionFiles(
	dataDir string,
	version containers.HskVersion,
	separator rune,
	processFields func (fields []string) (hskLevel containers.HskLevel, simplified string, traditional string, pinyinNum string, definition string),
	fileNames ...string) {

	hskWords, hskChars := containers.StringEntryMap{}, containers.CharEntryMap{}
	for _, fileName := range fileNames {
		parseHskFile(hskWords, hskChars, path.Join(dataDir, fileName), version, separator, processFields)
	}
	buildHskLevelToLists(hskWords, hskChars, version)
}



func parseHskFile(
  hskWords containers.StringEntryMap,
  hskChars containers.CharEntryMap,
  fileName string,
  version containers.HskVersion,
  separator rune,
  processFields func (fields []string) (hskLevel containers.HskLevel, simplified string, traditional string, pinyinNum string, definition string)) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = separator
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, each := range csvData {
		if strings.Contains(each[0], "\uFEFF") {
			each[0] = strings.Replace(each[0], "\uFEFF", "", -1)
		}
		hskLevel, simplified, traditional, pinyinNum, definition := processFields(each)
		processHskEntry(hskWords, hskChars, version, hskLevel, simplified, traditional, pinyinNum, definition)
	}
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
			var pinyinSyllable = ""
			if len(syllables) > 0 {
				pinyinSyllable = syllables[index]
			}
			var entry *containers.Entry
			if traditional == "" {
				entry = addOrUpdateEntry(string(char), "", pinyinSyllable, "", false)
			} else {
				entry = addOrUpdateEntry(string(char), string(traditionalChars[index]), pinyinSyllable, "", false)
			}
			entry.SetHskCharLevel(hskVersion, hskLevel)
			hskChars[char] = entry
		}
		entry := addOrUpdateEntry(simplified, traditional, cleanPinyin, definition, true)
		entry.SetHskWordLevel(hskVersion, hskLevel)
		hskWords[simplified] = entry
	}
}

func buildHskLevelToLists(hskWords containers.StringEntryMap, hskChars containers.CharEntryMap, hskVersion containers.HskVersion) {
	levelToWords := map[containers.HskLevel]containers.StringEntryMap{}
	levelToChars := map[containers.HskLevel]containers.CharEntryMap{}

	for i := containers.HskLevel(1); i <= HskVersionToLevels[hskVersion]; i++ {
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
