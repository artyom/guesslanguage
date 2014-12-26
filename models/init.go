package models

import (
	"sort"
	"unicode"
)

var models = make(map[string]map[Tg]int)

type Tg [3]rune

// compact stores string-defined trigrams with their weight into rune arrays
func compact(m map[string]int) map[Tg]int {
	out := make(map[Tg]int, len(m))
	for k, v := range m {
		var t Tg
		r := []rune(k)
		t = Tg{r[0], r[1], r[2]}
		out[t] = v
	}
	return out
}

// Struct used to sort trigrams
type valSorter struct {
	keys   []Tg
	values []int
}

// Returns list of all models
func GetModels() map[string]map[Tg]int {
	return models
}

// Create a list of trigrams in content sorted by frequency.
func GetOrderedModel(words [][]rune) []Tg {
	l := len(words) - 1
	for _, word := range words {
		l += len(word)
	}
	trigrams := make(map[Tg]int, l/3)
	var r0, r1 rune
	i := 0
	for _, word := range words {
		for _, r := range word {
			r = unicode.ToLower(r)
			t := Tg{r0, r1, r}
			r0, r1 = r1, r
			i++
			if i < 3 {
				continue
			}
			trigrams[t]++
		}
		if i < l {
			t := Tg{r0, r1, ' '}
			r0, r1 = r1, ' '
			i++
			if i < 3 {
				continue
			}
			trigrams[t]++
		}
	}

	vs := getValSorter(trigrams)
	vs.Sort()

	return vs.keys
}

func getValSorter(m map[Tg]int) *valSorter {
	vs := &valSorter{
		keys:   make([]Tg, 0, len(m)),
		values: make([]int, 0, len(m)),
	}

	for k, v := range m {
		vs.keys = append(vs.keys, k)
		vs.values = append(vs.values, v)
	}

	return vs
}

func (vs *valSorter) Sort() {
	sort.Sort(sort.Reverse(vs))
}

func (vs *valSorter) Len() int {
	return len(vs.values)
}

func (vs *valSorter) Less(i, j int) bool {
	return vs.values[i] < vs.values[j]
}

func (vs *valSorter) Swap(i, j int) {
	vs.values[i], vs.values[j] = vs.values[j], vs.values[i]
	vs.keys[i], vs.keys[j] = vs.keys[j], vs.keys[i]
}
