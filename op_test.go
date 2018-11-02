package quilldelta

import "testing"

func TestDeleteLength(t *testing.T) {
	ttc := []struct {
		op     Op
		length int
		desc   string
	}{
		{Op{Delete: newint(5)}, 5, `{delete:5}`},
		{Op{Retain: newint(2)}, 2, `{retain:2}`},
		{Op{Insert: "text"}, 4, `{insert:'text'}`},
		{Op{Insert: 2}, 1, `{insert:2}`},
	}

	for _, tc := range ttc {
		if Len(tc.op) != tc.length {
			t.Errorf("Expect length of %q to be %d, but got: %d", tc.desc, tc.length, Len(tc.op))
		}
	}
}
