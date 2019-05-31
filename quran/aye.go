package quran

import (
	"fmt"
	"math/rand"

	"github.com/globalsign/mgo/bson"
	"github.com/mavihq/persian"
	db "upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

// Aye strcut
type Aye struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	SuraID     bson.ObjectId `bson:"_sura_id,omitempty"`
	Number     uint          `bson:"number,omitempty"`
	Text       string        `bson:"text,omitempty"`
	SimpleText string        `bson:"simple_text,omitempty"`
	Translate  Translate     `bson:"translate"`
	Sura       Sura
}

// Translate struct
type Translate struct {
	FooladvandFa string `bson:"fa-fooladvand"`
	MakaremFa    string `bson:"fa-makarem"`
	GhomesheiFa  string `bson:"fa-ghomshei"`
}

// Sura struct
type Sura struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Number uint          `bson:"number,omitempty"`
	Name   string        `bson:"name,omitempty"`
	Ayat   uint          `bson:"cnt_aye"`
}

// newAyeByRand returns randomly an Aye from Quran
func newAyeByRand() (aye Aye, err error) {
	sess, err := mongo.Open(Mongo())
	if err != nil {
		return
	}
	defer sess.Close()

	res := sess.Collection(mongoDbAyeColl).Find().
		Limit(1).
		Offset(rand.Intn(6236))

	err = res.One(&aye)
	if err != nil {
		return
	}

	var sura Sura
	err = sess.Collection(mongoDbSuraColl).
		Find(db.Cond{"_id": aye.SuraID}).
		One(&sura)

	aye.Sura = sura

	return
}

// String prepare aye as string
func (a *Aye) String() string {
	return fmt.Sprintf("«%s»\n\n%s\n\n%s:%s",
		a.Text,
		a.Translate.FooladvandFa,
		a.Sura.Name,
		persian.ToPersianDigitsFromInt(int(a.Number)))
}
