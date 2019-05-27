package quran

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanTweet(t *testing.T) {
	// English string with 280 charachters
	assert.True(t, CanTweet("English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 char"))

	// Persian string with 280 charachters
	assert.True(t, CanTweet("متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن ف"))

	// English string with 281 charachters
	assert.False(t, CanTweet("English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 charachters English string with 280 chars"))

	// Persian string with 281 charachters
	assert.False(t, CanTweet("متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فارسی با ۲۸۰ کاراکتر متن فا"))
}
