package caesarcipher

// Encode encodes string s to caesar cipher by int shift
func Encode(s string, shift int) string {
	encoded := make([]byte, len(s))
	for i, c := range s {
		if c >= 97 && c <= 122 {
			n := int(c-97) + (shift % 26)
			for n > 26 {
				n -= 26
			}
			encoded[i] = alphabetsSmol[n]
		} else if c >= 65 && c <= 90 {
			n := int(c-65) + (shift % 26)
			for n > 26 {
				n -= 26
			}
			encoded[i] = alphabetsBig[n]
		} else {
			encoded[i] = byte(c)
		}
	}
	return string(encoded)
}
