package testdata

//go:generate go2elm --yml go2elm.yml --out out

type EmbeddedStruct struct {
	AString string  `json:"aString,omitempty"`
	AInt    int     `json:"aInt,omitempty"`
	AFloat  float64 `json:"aFloat,omitempty"`
	ABool   bool    `json:"aBool,omitempty"`
}

type Test1 struct {
	AString string  `json:"aString"`
	AInt    int     `json:"aInt"`
	AFloat  float64 `json:"aFloat"`
	ABool   bool    `json:"aBool"`
}

type Test2 struct {
	AStringPtr *string  `json:"aStringPtr"`
	AIntPtr    *int     `json:"aIntPtr"`
	AFloatPtr  *float64 `json:"aFloatPtr"`
	ABoolPtr   *bool    `json:"aBoolPtr"`
}

type Test3 struct {
	AStringSlice []string  `json:"aStringSlice"`
	AIntSlice    []int     `json:"aIntSlice"`
	AFloatSlice  []float64 `json:"aFloatSlice"`
	ABoolSlice   []bool    `json:"aBoolSlice"`
}

type Test4 struct {
	AStringObject struct{ X string }  `json:"aStringObject"`
	AIntObject    struct{ X int }     `json:"aIntObject"`
	AFloatObject  struct{ X float64 } `json:"aFloatObject"`
	ABoolObject   struct{ X bool }    `json:"aBoolObject"`
}

type Test5 struct {
	AStringMap map[string]string  `json:"aStringMap"`
	AIntMap    map[string]int     `json:"aIntMap"`
	AFloatMap  map[string]float64 `json:"aFloatMap"`
	ABoolMap   map[string]bool    `json:"aBoolMap"`

	EmbeddedStruct
}

type Test6 struct {
	EmbeddedStruct
}

type Test7 *Test2

type Test8 []*Test2

// TODO(shutej): Fix this error.
// type Test9 *Test9
