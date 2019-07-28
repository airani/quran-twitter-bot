package bot

import (
	"testing"
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
