package fp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	values := []string{"foo", "bar"}

	assert.Equal(t, true, Contains(values, "foo"))
}

func TestNotContains(t *testing.T) {
	values := []string{"foo", "bar"}

	assert.Equal(t, false, Contains(values, "foobar"))
}
