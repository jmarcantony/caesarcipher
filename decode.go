package caesercipher

func Decode(s string, shift int) []byte {
	decoded := make([]byte, len(s))
	for i, c := range s {
		if c >= 97 && c <= 122 {
			n := int(c-97) - (shift % 26)
			for n < 0 {
				n = 26 + n
			}
			decoded[i] = alphabetsSmol[n]
		} else if c >= 65 && c <= 90 {
			n := int(c-65) - (shift % 26)
			for n < 0 {
				n = 26 + n
			}
			decoded[i] = alphabetsBig[n]
		} else {
			decoded[i] = byte(c)
		}
	}
	return decoded
}
