package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapWhere(t *testing.T) {
	m := map[string]int{
		"Alice": 1000,
		"Bob":   2000,
		"Dave":  3000,
		"Greg":  4000,
	}
	f := func(kv KeyValue[string, int]) bool { return kv.Val < 2500 }

	res := MapWhere(m, f)

	assert.Equal(t, map[string]int{
		"Alice": 1000,
		"Bob":   2000,
	}, res)
}

func Test_MapSelect(t *testing.T) {
	m := map[string]int{
		"Alice": 10,
		"Bob":   20,
	}
	f := func(v int) int { return v * 2 }

	res := MapSelect(m, f)

	assert.Equal(t, map[string]int{
		"Alice": 20,
		"Bob":   40,
	}, res)
}

func Test_MapSingle_Success(t *testing.T) {
	m := map[string]int{
		"Alice": 10,
	}

	res, err := MapSingle(m)

	assert.Nil(t, err)
	assert.Equal(t, 10, res)
}

func Test_MapSingle_EmptySlice(t *testing.T) {
	m := map[string]int{}

	_, err := MapSingle(m)

	assert.NotNil(t, err)
}

func Test_MapSingle_NilSlice(t *testing.T) {
	var m map[string]int

	_, err := MapSingle(m)

	assert.NotNil(t, err)
}

func Test_MapSingle_SliceTooLong(t *testing.T) {
	m := map[string]int{
		"Alice": 10,
		"Bob":   20,
	}

	_, err := MapSingle(m)

	assert.NotNil(t, err)
}

func Test_MapAggregate(t *testing.T) {
	m := map[string]int{
		"Alice": 10,
		"Bob":   20,
	}
	f := func(a int, b int) int { return a + b }

	res := MapAggregate(m, 0, f)

	assert.Equal(t, 30, res)
}
