package base62Encoder

import "testing"

func TestEncode(t *testing.T) {
	var input = 12345
	var encoded = Encode(input)
	var decoded = Decode(encoded)
	if decoded != input {
		t.Errorf("Encode(%d) = %s, Decode(%s) = %d", input, encoded, encoded, decoded)
	}
}
