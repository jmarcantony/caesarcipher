package caesercipher

import "testing"

func TestDecode(t *testing.T) {
	cases := []struct {
		shift   int
		s, want string
	}{
		{1, "bcd", "abc"},
		{2, "okffng-Qwvb", "middle-Outz"},
	}
	for _, c := range cases {
		s := string(Decode(c.s, c.shift))
		if s != c.want {
			t.Errorf("got %q, want %q", s, c.want)
		}
	}
}
