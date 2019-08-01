package quran

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	datasetQuranSimpleFile         = "quran-simple-min.xml"
	datasetFooladvandTranslateFile = "fa.fooladvand.xml"
)

type datasetFile string

// Quran struct of xml file
type Quran struct {
	Suras []Sura `xml:"sura"`
}

var base Quran
var farsi Quran

func init() {
	var err error

	p := os.Getenv("PATH_DATASET")
	if p == "" {
		p = "../dataset/"
	}

	base, err = newQuranByXML(p + datasetQuranSimpleFile)
	if err != nil {
		log.Println(err.Error())
		return
	}

	farsi, err = newQuranByXML(p + datasetFooladvandTranslateFile)
	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}

// New Quran base instance
func New() Quran {
	return base
}

// Fa Quran Farsi translate instance
func Fa() Quran {
	return farsi
}

// newQuranByXML read xml file of quran and returns a Quran struct
func newQuranByXML(f string) (q Quran, err error) {
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
	if n > len(q.Suras) || n == 0 {
		return
	}
	return q.Suras[n-1]
}

// RandSura Returns a Sura by random
func (q Quran) RandSura() Sura {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(q.Suras)) + 1
	return q.Sura(n)
}
