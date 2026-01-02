package google_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/google"
	"testing"
)

func TestHasSameDigits(t *testing.T) {
	res := google.HasSameDigits("3902")
	assert.Equal(t, true, res)
}

func TestHasSameDigits2(t *testing.T) {
	res := google.HasSameDigits("34789")
	assert.Equal(t, false, res)
}

func TestHasSameDigits1(t *testing.T) {
	res := google.HasSameDigitsv2("3902")
	assert.Equal(t, true, res)
}

func TestHasSameDigits22(t *testing.T) {
	res := google.HasSameDigitsv2("34789")
	assert.Equal(t, false, res)
}

func TestHasSameDigits23(t *testing.T) {
	res := google.HasSameDigitsv2("2367679537183200808337444239103035091812101003664191823186564233")
	assert.Equal(t, true, res)
}
