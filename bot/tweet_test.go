package bot

import (
	"testing"

	"github.com/airani/quran"
)

func TestTweet_Valid(t1 *testing.T) {
	q, _ := quran.NewSimple()
	qf, _ := quran.NewTranslate(quran.TranslateFaFooladvand)
	type fields struct {
		Quran   quran.Quran
		QuranFa quran.Quran
	}
	qFields := fields{
		Quran:   q,
		QuranFa: qf,
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "English string with 280 characters",
			fields: qFields,
			args: args{
				s: "English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters.",
			},
			want: true,
		},
		{
			name:   "Persian string with 280 characters",
			fields: qFields,
			args: args{
				s: "متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن ف",
			},
			want: true,
		},
		{
			name:   "English string with 281 characters",
			fields: qFields,
			args: args{
				s: "English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters English string with 280 characters.x",
			},
			want: false,
		},
		{
			name:   "Persian string with 281 characters",
			fields: qFields,
			args: args{
				s: "متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فا",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tweet{
				Quran:   tt.fields.Quran,
				QuranFa: tt.fields.QuranFa,
			}
			if got := t.Valid(tt.args.s); got != tt.want {
				t1.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTweet_String(t1 *testing.T) {
	q, _ := quran.NewSimple()
	qf, _ := quran.NewTranslate(quran.TranslateFaFooladvand)
	type fields struct {
		Quran   quran.Quran
		QuranFa quran.Quran
	}
	qFields := fields{
		Quran:   q,
		QuranFa: qf,
	}
	type args struct {
		s quran.Surah
		a quran.Ayah
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			args: args{
				s: q.Surah(1),
				a: q.Surah(1).Ayah(1),
			},
			fields: qFields,
			want:   "«بِسمِ اللَّهِ الرَّحمنِ الرَّحيمِ»\n\nبه نام خداوند رحمتگر مهربان\n\nالفاتحة:۱",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tweet{
				Quran:   tt.fields.Quran,
				QuranFa: tt.fields.QuranFa,
			}
			if got := t.String(tt.args.s, tt.args.a); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
