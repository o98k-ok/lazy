package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestSearchStruct struct {
	Name string
	Age  int
}

func TestSearch(t *testing.T) {
	assert.Equal(t, []string{"1", "11"}, Search([]string{"1", "2", "3", "11"}, "1"))
	testcases := []TestSearchStruct{
		{"shadow", 10},
		{"tencent", 20},
		{"work", 100},
	}

	assert.Equal(t, []TestSearchStruct{{"shadow", 10}, {"work", 100}}, Search(testcases, "10"))
	assert.Equal(t, []TestSearchStruct{{"shadow", 10}, {"work", 100}}, Search(testcases, "o"))
}
