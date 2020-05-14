package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/airani/quran"
)

const (
	maxTweetLen int           = 280
	interval    time.Duration = 1 * time.Hour
)

// Run twitter bot
func Run() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var t Tweet
	var err error

	t.Quran, err = quran.NewSimple()
	if err != nil {
		log.Fatal(err.Error())
	}
	t.QuranFa, err = quran.NewTranslate(quran.TranslateFaFooladvand)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Quran twitter bot started..")

	ticker := time.NewTicker(interval)

	for range ticker.C {
		for {
			err := t.PostRandAyah()
			if err != nil {
				log.Println(fmt.Sprintf("%q", err))
			} else {
				break
			}
		}
	}
}
