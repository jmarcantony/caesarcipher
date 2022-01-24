package caesarcipher

import "testing"

func TestEncode(t *testing.T) {
	cases := []struct {
		shift   int
		s, want string
	}{
		{1, "abc", "bcd"},
		{2, "middle-Outz", "okffng-Qwvb"},
	}
	for _, c := range cases {
		s := Encode(c.s, c.shift)
		if s != c.want {
			t.Errorf("got %q, want %q", s, c.want)
		}
	}
}
