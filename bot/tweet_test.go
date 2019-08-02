package bot

import (
	"testing"

	"github.com/airani/quran-twitter-bot/quran"
)

func Test_canTweet(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "English string with 280 charachters",
			args: args{
				s: "English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 char",
			},
			want: true,
		},
		{
			name: "Persian string with 280 charachters",
			args: args{
				s: "متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن ف",
			},
			want: true,
		},
		{
			name: "English string with 281 charachters",
			args: args{
				s: "English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 chars",
			},
			want: false,
		},
		{
			name: "Persian string with 281 charachters",
			args: args{
				s: "متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فا",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canTweet(tt.args.s); got != tt.want {
				t.Errorf("canTweet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tweetAyeText(t *testing.T) {
	q := quran.New()

	type args struct {
		s quran.Sura
		a quran.Aya
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: q.Sura(1),
				a: q.Sura(1).Aya(1),
			},
			want: "«بِسمِ اللَّهِ الرَّحمنِ الرَّحيمِ»\n\nبه نام خداوند رحمتگر مهربان\n\nالفاتحة:۱",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tweetAyeText(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("tweetAyeText() = %v, want %v", got, tt.want)
			}
		})
	}
}
