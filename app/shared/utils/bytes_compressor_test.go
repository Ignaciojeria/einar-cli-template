package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Run with logs : go test -coverprofile=coverage.out ./app/shared/... -v
func TestStructToBytes(t *testing.T) {

	want := "string"

	encoded, _ := json.Marshal(want)

	compressed := CompressBytes(encoded)
	decompress, _ := DecompressBytes(compressed)

	var got string

	json.Unmarshal(decompress, &got)

	assert.Equal(t, want, got)
}
