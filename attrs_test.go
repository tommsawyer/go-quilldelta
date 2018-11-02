package quilldelta

import (
	"reflect"
	"testing"
)

func TestComposeAttributes(t *testing.T) {
	attrs := Attrs{
		"bold":  true,
		"color": "red",
	}

	tcs := []struct {
		msg      string
		a, b     Attrs
		expected Attrs
	}{
		{"When left is nil", nil, attrs, attrs},
		{"When right is nil", attrs, nil, attrs},
		{"When both are nil", nil, nil, Attrs{}},
		{"When attr missing", attrs, Attrs{"italic": true}, Attrs{"bold": true, "color": "red", "italic": true}},
		{"Overwriting attrs", attrs, Attrs{"bold": false, "color": "blue"}, Attrs{"bold": false, "color": "blue"}},
		{"Removing one attr", attrs, Attrs{"bold": nil}, Attrs{"color": "red"}},
		{"Removing all attrs", attrs, Attrs{"bold": nil, "color": nil}, Attrs{}},
		{"Removing missing", attrs, Attrs{"italic": nil}, attrs},
	}

	for _, tc := range tcs {
		composed := Compose(tc.a, tc.b)
		if !reflect.DeepEqual(composed, tc.expected) {
			t.Errorf("%s: composing %s with %s should return %s, but got: %s", tc.msg, tc.a, tc.b, tc.expected, composed)
		}
	}
}

func TestAttributesDiff(t *testing.T) {
	attrs := Attrs{
		"bold":  true,
		"color": "red",
	}

	tcs := []struct {
		msg      string
		a, b     Attrs
		expected Attrs
	}{
		{"When left is nil", nil, attrs, attrs},
		{"When right is nil", attrs, nil, Attrs{"bold": nil, "color": nil}},
		{"When both are the same", attrs, attrs, Attrs{}},
		{"Added param", attrs, Attrs{"bold": true, "color": "red", "italic": true}, Attrs{"italic": true}},
		{"Removed param", attrs, Attrs{"bold": true}, Attrs{"color": nil}},
		{"Overwrite param", attrs, Attrs{"bold": true, "color": "blue"}, Attrs{"color": "blue"}},
	}

	for _, tc := range tcs {
		diff := Diff(tc.a, tc.b)
		if !reflect.DeepEqual(diff, tc.expected) {
			t.Errorf("%s: diff between %s and %s should return %s, but got: %s", tc.msg, tc.a, tc.b, tc.expected, diff)
		}
	}
}
