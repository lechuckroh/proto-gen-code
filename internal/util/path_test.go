package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBaseFilename(t *testing.T) {
	result := GetBaseFilename("_test/constants_pb.ts")

	assert.Equal(t, "constants_pb", result)
}
