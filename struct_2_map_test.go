package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EmbendTest struct {
	EmbendInt    int    `mapKey:"embdMI"`
	EmbendString string `mapKey:"embdMS"`
}

type TestStruct struct {
	Int          int         `mapKey:"mi"`
	String       string      `mapKey:"ms"`
	Bool         bool        `mapKey:"mb"`
	IntList      []int       `mapKey:"ml"`
	StringList   []string    `mapKey:"ml2"`
	EmbenddedPtr *EmbendTest `mapKey:"mEmbdPtr"`
	Embendded    EmbendTest  `mapKey:"mEmbd"`
}

func Test_Stuct2Map(t *testing.T) {
	assert := assert.New(t)
	ori := TestStruct{
		Int:        1,
		String:     "string",
		Bool:       true,
		IntList:    []int{1, 2, 3},
		StringList: []string{"a", "b", "c"},
		EmbenddedPtr: &EmbendTest{
			EmbendInt:    10,
			EmbendString: "embendPTR",
		},
		Embendded: EmbendTest{
			EmbendInt:    20,
			EmbendString: "embend",
		},
	}
	structMap := AssignMap(ori, ParseStructType(ori, "mapKey"), true)
	assert.Equal(1, structMap["mi"])
	assert.Equal("string", structMap["ms"])
	assert.Equal(true, structMap["mb"])
	assert.Equal([]int{1, 2, 3}, structMap["ml"])
	assert.Equal([]string{"a", "b", "c"}, structMap["ml2"])
	m := structMap["mEmbdPtr"].(map[string]interface{})
	assert.Equal(10, m["embdMI"])
	assert.Equal("embendPTR", m["embdMS"])
	m2 := structMap["mEmbd"].(map[string]interface{})
	assert.Equal(20, m2["embdMI"])
	assert.Equal("embend", m2["embdMS"])
}
