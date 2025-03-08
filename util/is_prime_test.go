package util_test

import (
	"github.com/stretchr/testify/assert"
	"gocoder/util"
	"testing"
)

func TestIsPrime(t *testing.T) {
	t.Run("test primes", func(t *testing.T) {
		assert.Equal(t, false, util.IsPrime(1))
		assert.Equal(t, true, util.IsPrime(2))
		assert.Equal(t, true, util.IsPrime(3))
		assert.Equal(t, false, util.IsPrime(4))
		assert.Equal(t, true, util.IsPrime(5))
		assert.Equal(t, true, util.IsPrime(7))
		assert.Equal(t, false, util.IsPrime(9))
		assert.Equal(t, true, util.IsPrime(11))
		assert.Equal(t, true, util.IsPrime(13))
		assert.Equal(t, false, util.IsPrime(15))
	})
}
