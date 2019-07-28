package quran

import (
	"encoding/xml"
	"io/ioutil"
	"math/rand"
	"os"
)

const (
	datasetQuranSimpleFile         datasetFile = "dataset/quran-simple-min.xml"
	datasetFooladvandTranslateFile             = "dataset/fa.fooladvand.xml"
)

type datasetFile string

// Quran struct of xml file
type Quran struct {
	Suras []Sura `xml:"sura"`
}

// M is Main struct of Quran
var M Quran

// Fa is Farsi translate of Quran
var Fa Quran

func init() {
	var err error

	M, err = newQuranByXML(datasetQuranSimpleFile)
	if err != nil {
		return
	}

	Fa, err = newQuranByXML(datasetFooladvandTranslateFile)
	if err != nil {
		return
	}

	return
}

// newQuranByXML read xml file of quran and returns a Quran struct
func newQuranByXML(f datasetFile) (q Quran, err error) {
	xmlFile, err := os.Open(string(f))
	if err != nil {
		return
	}

	defer xmlFile.Close()

	b, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return
	}

	xml.Unmarshal(b, &q)

	return
}

// Sura Returns a Sura by number
func (q Quran) Sura(n int) (s Sura) {
	if n > len(q.Suras) {
		return
	}
	return q.Suras[n]
}

// RandSura Returns a Sura by random
func (q Quran) RandSura() Sura {
	n := rand.Intn(len(q.Suras)-1) + 1
	return q.Sura(n)
}
