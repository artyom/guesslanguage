// Package guesslanguage provides a simple way to detect language of unicode string.
//
// Source code and project home:
// https://github.com/endeveit/guesslanguage
//
package guesslanguage

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/endeveit/guesslanguage/models"
)

func init() {
	// put unicode.Scripts map values into slice as iterating over slice is
	// faster than over map. TODO: as further optimization, slice can be
	// sorted to place most expected or frequently used scripts at the
	// beginning
	knownScripts = make([]*unicode.RangeTable, 0, len(unicode.Scripts))
	knownScripts = append(knownScripts, unicode.Latin, unicode.Cyrillic)
	for _, s := range unicode.Scripts {
		if s == unicode.Latin || s == unicode.Cyrillic {
			continue
		}
		knownScripts = append(knownScripts, s)
	}
}

var (
	knownScripts    []*unicode.RangeTable
	maxLength       int      = 4096
	maxDistance     int      = maxLength * 300
	minLength       int      = 20
	maxGrams        int      = 300
	unknownLanguage string   = "UNKNOWN"
	codesBasicLatin []string = []string{
		"ceb",
		"en",
		"eu",
		"ha",
		"haw",
		"id",
		"la",
		"nr",
		"nso",
		"so",
		"ss",
		"st",
		"sw",
		"tlh",
		"tn",
		"ts",
		"xh",
		"zu"}
	codesExtendedLatin []string = []string{
		"af",
		"az",
		"ca",
		"cs",
		"cy",
		"da",
		"de",
		"eo",
		"es",
		"et",
		"fi",
		"fr",
		"hr",
		"hu",
		"is",
		"it",
		"lt",
		"lv",
		"nb",
		"nl",
		"pl",
		"pt",
		"ro",
		"sk",
		"sl",
		"sq",
		"sv",
		"tl",
		"tr",
		"ve",
		"vi"}
	codesAllLatin   []string = append(codesBasicLatin, codesExtendedLatin...)
	codesCyrillic   []string = []string{"bg", "kk", "ky", "mk", "mn", "ru", "sr", "uk", "uz"}
	codesArabic     []string = []string{"ar", "fa", "ps", "ur"}
	codesDevanagari []string = []string{"hi", "ne"}
	codesPt         []string = []string{"pt_BR", "pt_PT"}
	// NOTE mn appears twice, once for mongolian script and once for codesCyrillic
	singletons map[string]string = map[string]string{
		"Armenian":  "hy",
		"Hebrew":    "he",
		"Bengali":   "bn",
		"Gurmukhi":  "pa",
		"Greek":     "el",
		"Gujarati":  "gu",
		"Oriya":     "or",
		"Tamil":     "ta",
		"Telugu":    "te",
		"Kannada":   "kn",
		"Malayalam": "ml",
		"Sinhala":   "si",
		"Thai":      "th",
		"Lao":       "lo",
		"Tibetan":   "bo",
		"Burmese":   "my",
		"Georgian":  "ka",
		"Mongolian": "mn-Mong",
		"Khmer":     "km"}
	nameMap map[string]string = map[string]string{
		"ab":    "Abkhazian",
		"af":    "Afrikaans",
		"ar":    "Arabic",
		"az":    "Azeri",
		"be":    "Byelorussian",
		"bg":    "Bulgarian",
		"bn":    "Bengali",
		"bo":    "Tibetan",
		"br":    "Breton",
		"ca":    "Catalan",
		"ceb":   "Cebuano",
		"cs":    "Czech",
		"cy":    "Welsh",
		"da":    "Danish",
		"de":    "German",
		"el":    "Greek",
		"en":    "English",
		"eo":    "Esperanto",
		"es":    "Spanish",
		"et":    "Estonian",
		"eu":    "Basque",
		"fa":    "Farsi",
		"fi":    "Finnish",
		"fo":    "Faroese",
		"fr":    "French",
		"fy":    "Frisian",
		"gd":    "Scots Gaelic",
		"gl":    "Galician",
		"gu":    "Gujarati",
		"ha":    "Hausa",
		"haw":   "Hawaiian",
		"he":    "Hebrew",
		"hi":    "Hindi",
		"hr":    "Croatian",
		"hu":    "Hungarian",
		"hy":    "Armenian",
		"id":    "Indonesian",
		"is":    "Icelandic",
		"it":    "Italian",
		"ja":    "Japanese",
		"ka":    "Georgian",
		"kk":    "Kazakh",
		"km":    "Cambodian",
		"ko":    "Korean",
		"ku":    "Kurdish",
		"ky":    "Kyrgyz",
		"la":    "Latin",
		"lt":    "Lithuanian",
		"lv":    "Latvian",
		"mg":    "Malagasy",
		"mk":    "Macedonian",
		"ml":    "Malayalam",
		"mn":    "Mongolian",
		"mr":    "Marathi",
		"ms":    "Malay",
		"nd":    "Ndebele",
		"ne":    "Nepali",
		"nl":    "Dutch",
		"nn":    "Nynorsk",
		"no":    "Norwegian",
		"nso":   "Sepedi",
		"pa":    "Punjabi",
		"pl":    "Polish",
		"ps":    "Pashto",
		"pt":    "Portuguese",
		"pt_PT": "Portuguese (Portugal)",
		"pt_BR": "Portuguese (Brazil)",
		"ro":    "Romanian",
		"ru":    "Russian",
		"sa":    "Sanskrit",
		"sh":    "Serbo-Croatian",
		"sk":    "Slovak",
		"sl":    "Slovene",
		"so":    "Somali",
		"sq":    "Albanian",
		"sr":    "Serbian",
		"sv":    "Swedish",
		"sw":    "Swahili",
		"ta":    "Tamil",
		"te":    "Telugu",
		"th":    "Thai",
		"tl":    "Tagalog",
		"tlh":   "Klingon",
		"tn":    "Setswana",
		"tr":    "Turkish",
		"ts":    "Tsonga",
		"tw":    "Twi",
		"uk":    "Ukrainian",
		"ur":    "Urdu",
		"uz":    "Uzbek",
		"ve":    "Venda",
		"vi":    "Vietnamese",
		"xh":    "Xhosa",
		"zh":    "Chinese",
		"zh_TW": "Traditional Chinese (Taiwan)",
		"zu":    "Zulu"}
	ianaMap map[string]int = map[string]int{
		"ab":    12026,
		"af":    40,
		"ar":    26020,
		"az":    26030,
		"be":    11890,
		"bg":    26050,
		"bn":    26040,
		"bo":    26601,
		"br":    1361,
		"ca":    3,
		"ceb":   26060,
		"cs":    26080,
		"cy":    26560,
		"da":    26090,
		"de":    26160,
		"el":    26165,
		"en":    26110,
		"eo":    11933,
		"es":    26460,
		"et":    26120,
		"eu":    1232,
		"fa":    26130,
		"fi":    26140,
		"fo":    11817,
		"fr":    26150,
		"fy":    1353,
		"gd":    65555,
		"gl":    1252,
		"gu":    26599,
		"ha":    26170,
		"haw":   26180,
		"he":    26592,
		"hi":    26190,
		"hr":    26070,
		"hu":    26200,
		"hy":    26597,
		"id":    26220,
		"is":    26210,
		"it":    26230,
		"ja":    26235,
		"ka":    26600,
		"kk":    26240,
		"km":    1222,
		"ko":    26255,
		"ku":    11815,
		"ky":    26260,
		"la":    26280,
		"lt":    26300,
		"lv":    26290,
		"mg":    1362,
		"mk":    26310,
		"ml":    26598,
		"mn":    26320,
		"mr":    1201,
		"ms":    1147,
		"ne":    26330,
		"nl":    26100,
		"nn":    172,
		"no":    26340,
		"pa":    65550,
		"pl":    26380,
		"ps":    26350,
		"pt":    26390,
		"ro":    26400,
		"ru":    26410,
		"sa":    1500,
		"sh":    1399,
		"sk":    26430,
		"sl":    26440,
		"so":    26450,
		"sq":    26010,
		"sr":    26420,
		"sv":    26480,
		"sw":    26470,
		"ta":    26595,
		"te":    26596,
		"th":    26594,
		"tl":    26490,
		"tlh":   26250,
		"tn":    65578,
		"tr":    26500,
		"tw":    1499,
		"uk":    26520,
		"ur":    26530,
		"uz":    26540,
		"vi":    26550,
		"zh":    26065,
		"zh_TW": 22}
)

// Return the ISO 639-1 language code.
func Guess(text string) (result string, err error) {
	if !utf8.ValidString(text) {
		return result, errors.New("Input string contains invalid UTF-8-encoded runes")
	}
	words := make([][]rune, 0)
	wrd := make([]rune, 0, 10)
	cnt := 0
runeExtract:
	for _, r := range text {
		switch {
		case cnt > maxLength:
			break runeExtract
		case r == '’':
			wrd = append(wrd, '\'')
		case unicode.IsLetter(r):
			wrd = append(wrd, r)
		case unicode.IsSpace(r):
			words = append(words, wrd)
			wrd = make([]rune, 0, 10)
		default:
			continue
		}
		cnt++
	}
	if len(wrd) > 0 {
		words = append(words, wrd)
	}

	return guessLanguage(words, getRuns(words)), nil
}

// Return the language IANA ID.
func GuessId(text string) int {
	code, err := Guess(text)

	if err != nil || code == unknownLanguage {
		return 0
	}

	return ianaMap[code]
}

// Return the language name (in English).
func GuessName(text string) string {
	code, err := Guess(text)

	if err != nil || code == unknownLanguage {
		return unknownLanguage
	}

	return nameMap[code]
}

// Identify the language.
func guessLanguage(words [][]rune, scripts []*unicode.RangeTable) string {
	if keyExists(unicode.Hangul, scripts) {
		return "ko"
	}

	if keyExists(unicode.Greek, scripts) || keyExists(unicode.Coptic, scripts) {
		return "el"
	}

	if keyExists(unicode.Hiragana, scripts) ||
		keyExists(unicode.Katakana, scripts) {
		return "ja"
	}

	if keyExists(unicode.Han, scripts) ||
		keyExists(unicode.Bopomofo, scripts) {
		return "zh"
	}

	if keyExists(unicode.Cyrillic, scripts) {
		return getFromModel(words, codesCyrillic)
	}

	if keyExists(unicode.Arabic, scripts) {
		return getFromModel(words, codesArabic)
	}

	if keyExists(unicode.Devanagari, scripts) {
		return getFromModel(words, codesDevanagari)
	}

	// Try languages with unique scripts
	for blockName, langName := range singletons {
		if keyExists(unicode.Scripts[blockName], scripts) {
			return langName
		}
	}

	if keyExists(unicode.Latin, scripts) {
		latinLang := getFromModel(words, codesAllLatin)
		if latinLang == "pt" {
			return getFromModel(words, codesPt)
		}
		return latinLang
	}

	return unknownLanguage
}

// Count the number of characters in each character block
func getRuns(words [][]rune) (relevantRuns []*unicode.RangeTable) {
	type scriptFreq struct {
		block *unicode.RangeTable
		cnt   int
	}
	var (
		runTypes     = make([]*scriptFreq, 0, 1)
		nbTotalChars = 0
		charBlock    *unicode.RangeTable
		percentage   int
		found        bool
	)

	for _, word := range words {
		for _, char := range word {
			charBlock = nil
			for _, s := range knownScripts {
				if unicode.Is(s, char) {
					charBlock = s
					break
				}
			}
			if charBlock == nil {
				continue
			}
			found = false
			for _, item := range runTypes {
				if item.block == charBlock {
					item.cnt++
					found = true
					break
				}
			}
			if !found {
				runTypes = append(runTypes, &scriptFreq{block: charBlock, cnt: 1})
			}
			nbTotalChars++
		}
	}

	// return run types that used for 40% or more of the string
	// return Basic Latin if found more than 15%
	for _, item := range runTypes {
		percentage = item.cnt * 100

		if percentage >= 40 || percentage >= 15 && item.block == unicode.Latin {
			relevantRuns = append(relevantRuns, item.block)
		}
	}

	return relevantRuns
}

// Check if key exists
func keyExists(a *unicode.RangeTable, list []*unicode.RangeTable) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Check words against known models
func getFromModel(words [][]rune, languages []string) (result string) {
	l := len(words) - 1
	for _, word := range words {
		l += len(word)
	}
	if l < minLength {
		return unknownLanguage
	}

	var (
		scores  map[string]int = make(map[string]int, len(languages))
		model                  = models.GetOrderedModel(words)
		minimal int            = maxDistance
	)

	for _, lang := range languages {
		scores[lang] = getDistance(model, models.GetModels()[strings.ToLower(lang)])
	}

	for lang, score := range scores {
		if score < minimal {
			minimal = score
			result = lang
		}
	}

	return result
}

// Calculate the distance to the known model.
func getDistance(model []models.Tg, known_model map[models.Tg]int) int {
	var (
		data []models.Tg
		dist int
		sub  int
	)

	if len(model) > maxGrams {
		data = model[:maxGrams]
	} else {
		data = model
	}

	for i, value := range data {
		if _, ok := known_model[value]; ok {
			sub = i - known_model[value]
			if sub < 0 {
				sub = -sub
			}

			dist += sub
		} else {
			dist += maxGrams
		}
	}

	return dist
}
