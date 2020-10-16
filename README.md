# struct_to_map

Change a strut into map[string]interface{}, we can set map with original struct param name or with name setted in tag

```

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
```
structMap will be assigned as following:
```
{
	"mEmbd": {
		"embdMI": 20,
		"embdMS": "embend"
	},
	"mEmbdPtr": {
		"embdMI": 10,
		"embdMS": "embendPTR"
	},
	"mb": true,
	"mi": 1,
	"ml": [
		1,
		2,
		3
	],
	"ml2": [
		"a",
		"b",
		"c"
	],
	"ms": "string"
}
```

