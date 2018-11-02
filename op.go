package quilldelta

import "unicode/utf8"

type Op struct {
	Insert interface{} `json:"insert,omitempty"`
	Delete *int        `json:"delete,omitempty"`
	Retain *int        `json:"retain,omitempty"`
	Attrs  Attrs       `json:"attributes,omitempty"`
}

func newint(n int) *int          { return &n }
func newstring(s string) *string { return &s }

func Len(op Op) int {
	if op.Delete != nil {
		return *op.Delete
	}

	if op.Retain != nil {
		return *op.Retain
	}

	if insertStr, ok := op.Insert.(string); ok {
		return utf8.RuneCountInString(insertStr)
	}

	return 1
}
