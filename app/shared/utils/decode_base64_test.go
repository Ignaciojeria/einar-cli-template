package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBase64(t *testing.T) {
	t.Parallel()

	t.Run("Success DecodeBase64", func(t *testing.T) {

		want := []byte("that is a string decode")

		got, _ := DecodeBase64("dGhhdCBpcyBhIHN0cmluZyBkZWNvZGU=")

		assert.Equal(t, want, got)

	})

	t.Run("Error in DecodeBase64", func(t *testing.T) {

		_, err := DecodeBase64("error")

		assert.Error(t, err)

	})
}
