package quran

import "math/rand"

// Sura struct of a sura of Quran
type Sura struct {
	Index int    `xml:"index,attr"`
	Name  string `xml:"name,attr"`
	Ayas  []Aya  `xml:"aya"`
}

// Aya Returns a Aya by number
func (s Sura) Aya(n int) (a Aya) {
	if n > len(s.Ayas) {
		return
	}
	return s.Ayas[n]
}

// RandAya Returns a Aya by random
func (s Sura) RandAya() Aya {
	n := rand.Intn(len(s.Ayas)-1) + 1
	return s.Aya(n)
}
