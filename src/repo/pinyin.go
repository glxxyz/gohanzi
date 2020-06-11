package repo

import "regexp"

// based on: https://stackoverflow.com/questions/20736291/regex-for-matching-pinyin
var pinyinSyllable =
	`([mM]iu|[pmPM]ou|[bpmBPM](o|e(i|ng?)?|a(ng?|i|o)?|i(e|ng?|a[no])?|u))|` +
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

var pinyinSyllableRegEx = regexp.MustCompile(`'?((` + pinyinSyllable + `)[0-9]?)\s*`)
var pinyinTonelessRegEx = regexp.MustCompile(`'?(` + pinyinSyllable + `)[0-9]?\s*`)

var pinyinSyllableRegEx2 = regexp.MustCompile(
	`(` +
  `([bpmfdtnlgkhjqxzcsrwy]|zh|ch|sh)?` +
  `([aeiou]|ai|an|ang|ao|ei|en|eng|er|ia|ian|iang|iao|ie|in|ing|iong|iu|ong|ou|ua|uai|uan|uang|ue|ui|un|un|uo)` +
  `[0-9]?` +
  `)\s*`)

var pinyinTonelessRegEx2 = regexp.MustCompile(
	`(` +
  `([bpmfdtnlgkhjqxzcsrwy]|zh|ch|sh)?` +
  `([aeiou]|ai|an|ang|ao|ei|en|eng|er|ia|ian|iang|iao|ie|in|ing|iong|iu|ong|ou|ua|uai|uan|uang|ue|ui|un|un|uo)` +
  `)[0-9]?\s*`)

func pinyinToSyllables(pinyin string) (syllables []string) {
	result := make([]string, 0)
	var matches = pinyinSyllableRegEx.FindAllStringSubmatch(pinyin, -1)
	for _, match := range matches {
		result = append(result, match[1])
	}
	return result
}

func pinyinTonelessSyllables(pinyin string) (syllables []string) {
	result := make([]string, 0)
	var matches = pinyinTonelessRegEx.FindAllStringSubmatch(pinyin, -1)
	for _, match := range matches {
		result = append(result, match[1])
	}
	return result
}
