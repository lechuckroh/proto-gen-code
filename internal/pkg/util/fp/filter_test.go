package fp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isEven := func(n int) bool { return n%2 == 0 }

	expected := []int{2, 4, 6, 8, 10}
	actual := Filter(numbers, isEven)
	assert.Equal(t, expected, actual)
}
