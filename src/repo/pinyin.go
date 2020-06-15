package repo

import (
	"regexp"
	"strings"
)

// based on: https://stackoverflow.com/questions/20736291/regex-for-matching-pinyin
const pinyinSyllable = `([mM]iu|[pmPM]ou|[bpmBPM](o|e(i|ng?)?|a(ng?|i|o)?|i(e|ng?|a[no])?|u))|` +
  `([fF](ou?|[ae](ng?|i)?|u))|([dD](e(i|ng?)|i(a[on]?|u))|` +
  `[dtDT](a(i|ng?|o)?|e(i|ng)?|i(a[on]?|e|ng|u)?|o(ng?|u)|u(o|i|an?|n)?))|` +
  `([nN]eng?|[lnLN](a(i|ng?|o)?|e(i|ng)?|i(ang|a[on]?|e|ng?|u)?|o(ng?|u)|u(o|i|an?|n)?|ve?))|` +
  `([ghkGHK](a(i|ng?|o)?|e(i|ng?)?|o(u|ng)|u(a(i|ng?)?|i|n|o)?))|` +
  `([zZ]h?ei|[czCZ]h?(e(ng?)?|o(ng?|u)?|ao|u?a(i|ng?)?|u?(o|i|n)?))|` +
  `([sS]ong|[sS]hua(i|ng?)?|[sS]hei|[sS][h]?(a(i|ng?|o)?|en?g?|ou|u(a?n|o|i)?|i))|` +
  `([rR]([ae]ng?|i|e|ao|ou|ong|u[oin]|ua?n?))|` +
  `([jqxJQX](i(a(o|ng?)?|[eu]|ong|ng?)?|u(e|a?n)?))|` +
  `(([aA](i|o|ng?)?|[oO]u?|[eE](i|ng?|r)?))|` +
  `([wW](a(i|ng?)?|o|e(i|ng?)?|u))|` +
  `[yY](a(o|ng?)?|e|in?g?|o(u|ng)?|u(e|a?n)?)|` +
  `r|ng|fun|v`

var pinyinSyllableRegEx = regexp.MustCompile(`\s*'?((` + pinyinSyllable + `)[1-4]?)5?\s*`)
var pinyinTonelessRegEx = regexp.MustCompile(`\s*'?(` + pinyinSyllable + `)[1-5]?\s*`)

func parsePinyinNumTones(pinyin string) (syllables []string, cleanPinyin string) {
	return syllablesFromRegex(pinyin, pinyinSyllableRegEx)
}

func parsePinyinNumToneless(pinyin string) (syllables []string, cleanPinyin string) {
	return syllablesFromRegex(pinyin, pinyinTonelessRegEx)
}

func syllablesFromRegex(pinyin string, regex *regexp.Regexp) (syllables []string, cleanPinyin string) {
	pinyinFixed := strings.Replace(pinyin, "ü", "v", -1)
	pinyinFixed = strings.Replace(pinyinFixed, "Ü", "V", -1)
	result := make([]string, 0)
	matches := regex.FindAllStringSubmatch(pinyinFixed, -1)
	for _, match := range matches {
		result = append(result, match[1])
	}
	return result, strings.Join(result, " ")
}
