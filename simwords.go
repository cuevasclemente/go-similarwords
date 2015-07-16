// A simple jaro-winkler pacakge
// that defines a jaro-winkler distance
// to be between 0 and 1
// (a port of SimilarWords.jl)
package gojw

import "strings"

func OrderStringsByLength(s1, s2 string) (shorter, longer string) {
	if len(s1) > len(s2) {
		return s2, s1
	}
	return s1, s2
}

func StringContains(s1, s2 string) (inner, outer string) {
	Inner, Outer := OrderStringsByLength(s1, s2)
	inner = strings.ToLower(Inner)
	outer = strings.ToLower(Outer)
	return
}
func SharedWords(inner, outer string) (wordsContained int) {
	iWords := getWords(inner)
	oWords := getWords(outer)
	for iWord := range iWords {
		for oWord := range oWords {
			if iWord == oWord {
				wordsContained++
			}
		}
	}
	return
}

func getWords(s string) (m map[string]bool) {
	m = make(map[string]bool)
	words := strings.Split(s, " ")
	for _, word := range words {
		m[word] = true
	}
	return
}

func JaroWinkler(s1, s2 string) (m float64) {
	inner, outer := OrderStringsByLength(s1, s2)
	l := float64(len(getWords(inner)))
	shared := float64(SharedWords(inner, outer))
	m = 1.0 - (shared / l)
	return
}
