package quilldelta

import (
	"encoding/json"
	"reflect"
)

type Attrs map[string]interface{}

func (a Attrs) String() string {
	bs, _ := json.Marshal(a)
	return string(bs)
}

func Compose(a, b Attrs) Attrs {
	result := Attrs{}

	if len(a) == 0 && len(b) == 0 {
		return result
	}

	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	for k, v := range a {
		result[k] = v
	}

	for k, v := range b {
		if v == nil {
			delete(result, k)
			continue
		}

		result[k] = v
	}

	return result
}

func Diff(a, b Attrs) Attrs {
	result := Attrs{}

	keys := make([]string, 0, len(a)+len(b))
	for k := range a {
		keys = append(keys, k)
	}
	for k := range b {
		keys = append(keys, k)
	}

	for _, key := range keys {
		if !reflect.DeepEqual(a[key], b[key]) {
			result[key] = b[key]
		}
	}

	return result
}
