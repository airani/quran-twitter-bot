package quran

const (
	// SajdaRecommended Sajda for aya recommended
	SajdaRecommended = "recommended"
	// SajdaObligatory Sajda for aya obligatory
	SajdaObligatory = "obligatory"
)

// Aya returns Aya struct of Sura
type Aya struct {
	Index     int    `xml:"index,attr"`
	Text      string `xml:"text,attr"`
	Bismillah string `xml:"bismillah,attr"`
	Sajda     string `xml:"sajda,attr,omitempty"`
}

// IsSajdaObligatory returns Aya is sajda obligatory
func (a Aya) IsSajdaObligatory() bool {
	if a.Sajda == SajdaObligatory {
		return true
	}
	return false
}
