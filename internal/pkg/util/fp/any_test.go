package fp

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAnyTrue(t *testing.T) {
	values := []string{"foo", "bar"}

	assert.Equal(t, true, Any(values, func(s string) bool {
		return strings.HasPrefix(s, "foo")
	}))
}

func TestAnyFalse(t *testing.T) {
	values := []string{"foo", "bar"}

	assert.Equal(t, false, Any(values, func(s string) bool {
		return strings.HasPrefix(s, "foobar")
	}))
}
