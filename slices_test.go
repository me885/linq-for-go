package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceWhere(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(i int) bool { return i%2 == 0 }

	res := SliceWhere(s, f)

	assert.Equal(t, []int{2, 4}, res)
}

func Test_SliceSelect(t *testing.T) {
	s := []int{1, 2, 3, 4}
	f := func(x int) int { return x * 2 }

	res := SliceSelect(s, f)

	assert.Equal(t, []int{2, 4, 6, 8}, res)
}

func Test_SliceFirst_Success(t *testing.T) {
	s := []int{1, 2, 3}

	res, err := SliceFirst(s)

	assert.Nil(t, err)
	assert.Equal(t, 1, res)
}

func Test_SliceFirst_EmptySlice(t *testing.T) {
	s := []int{}

	_, err := SliceFirst(s)

	assert.NotNil(t, err)
}

func Test_SliceFirst_NilSlice(t *testing.T) {
	var s []int

	_, err := SliceFirst(s)

	assert.NotNil(t, err)
}

func Test_SliceSingle_Success(t *testing.T) {
	s := []int{3}

	res, err := SliceSingle(s)

	assert.Nil(t, err)
	assert.Equal(t, 3, res)
}

func Test_SliceSingle_EmptySlice(t *testing.T) {
	s := []int{}

	_, err := SliceSingle(s)

	assert.NotNil(t, err)
}

func Test_SliceSingle_NilSlice(t *testing.T) {
	var s []int

	_, err := SliceSingle(s)

	assert.NotNil(t, err)
}

func Test_SliceSingle_SliceTooLong(t *testing.T) {
	s := []int{1, 2, 3}

	_, err := SliceSingle(s)

	assert.NotNil(t, err)
}
