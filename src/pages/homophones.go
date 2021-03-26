package pages

import (
	"fmt"
	"github.com/glxxyz/gohanzi/containers"
	"github.com/glxxyz/gohanzi/repo"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HomophonesParams struct {
	Expand     bool
	NumChars   int8
	MatchTones bool
	HskVersion int
	Homophones []repo.Homophone
}

func HomophonesHandler(response http.ResponseWriter, request *http.Request, start time.Time) {
	params := HomophonesParams{
		Expand:     request.FormValue("Expand") == "yes",
		NumChars:   formValueInt8(request, "chars", 2),
		MatchTones: request.FormValue("tones") == "yes",
		HskVersion: formValueInt(request,"hskVersion", 0),
	}
	funcs := template.FuncMap{
		"homophonesLink":   homophonesLink(params),
		"dictionaryLink":   dictionaryLink,
		"pinyinSearchLink": pinyinSearchLink(params),
	}
	hskVersion := parseHskVersion(params.HskVersion)
	params.Homophones = repo.BuildHomophones(params.NumChars, params.MatchTones, hskVersion)
	if err := executeTemplate(response, start, "homophones.gohtml", params, funcs); err != nil {
		panic(err)
	}
}

func parseHskVersion(version int) containers.HskVersion {
	switch version {
	case 1992:
		return repo.Hsk1992
	case 2010:
		return repo.Hsk2010
	case 2012:
		return repo.Hsk2012
	case 2020:
		return repo.Hsk2020
	}
	return repo.HskNone
}

func homophonesLink(params HomophonesParams) func(change string) string {
	return func(change string) string {
		expand := params.Expand
		numChars := params.NumChars
		hskVersion := params.HskVersion
		matchTones := params.MatchTones
		switch change {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			numCharsInt, _ := strconv.Atoi(change)
			numChars = int8(numCharsInt)
		case "invertExpand":
			expand = !expand
		case "invertTones":
			matchTones = !matchTones
		case "allWords":
			hskVersion = 0
		case "hsk1992":
			hskVersion = 1992
		case "hsk2010":
			hskVersion = 2010
		case "hsk2012":
			hskVersion = 2012
		case "hsk2020":
			hskVersion = 2020
		}
		request, err := http.NewRequest("GET", "/homophones", nil)
		if err != nil {
			panic(err)
		}
		query := request.URL.Query()
		query.Add("chars", fmt.Sprintf("%v", numChars))
		if expand {
			query.Add("expand", "yes")
		}
		if matchTones {
			query.Add("tones", "yes")
		}
		query.Add("hskVersion", fmt.Sprintf("%v", hskVersion))
		request.URL.RawQuery = query.Encode()
		return request.URL.String()
	}
}

func pinyinSearchLink(params HomophonesParams) func(pinyin string) string {
	return func(pinyin string) string {
		request, err := http.NewRequest("GET", "/search", nil)
		if err != nil {
			panic(err)
		}
		query := request.URL.Query()
		query.Add("format", "regex")
		if params.MatchTones {
			query.Add("pinyin", pinyin)
		} else {
			query.Add("pinyin", strings.Replace(pinyin, " ", "\\d?", -1))
		}
		if params.HskVersion != 0 {
			query.Add("hskVersion", string(params.HskVersion))
		}
		request.URL.RawQuery = query.Encode()
		return request.URL.String()
	}
}
