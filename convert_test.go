package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceToMap(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	slice := []Person{
		{Name: "Bill", Age: 34},
		{Name: "Bob", Age: 23},
	}

	key := func(_ int, p Person) string {
		return p.Name
	}

	res := SliceToMap(slice, key)

	assert.Equal(t, map[string]Person{
		"Bill": {Name: "Bill", Age: 34},
		"Bob":  {Name: "Bob", Age: 23},
	}, res)
}

func Test_MapToSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	m := map[string]Person{
		"Bill": {Name: "Bill", Age: 34},
		"Bob":  {Name: "Bob", Age: 23},
	}

	res := MapToSlice(m)

	assert.Equal(t, []Person{
		{Name: "Bill", Age: 34},
		{Name: "Bob", Age: 23},
	}, res)
}

func Test_MapSelectToSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	m := map[string]int{
		"Bill": 34,
		"Bob":  23,
	}

	f := func(kv KeyValue[string, int]) Person {
		return Person{Name: kv.Key, Age: kv.Val}
	}

	res := MapSelectToSlice(m, f)

	assert.Equal(t, []Person{
		{Name: "Bill", Age: 34},
		{Name: "Bob", Age: 23},
	}, res)
}

func Test_MapKeyToSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	m := map[string]Person{
		"Bill": {Name: "Bill", Age: 34},
		"Bob":  {Name: "Bob", Age: 23},
	}

	res := MapKeysToSlice(m)

	assert.Equal(t, []string{"Bill", "Bob"}, res)
}
