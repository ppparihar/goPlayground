package base62Encoder

import "testing"

func TestEncode(t *testing.T) {
	var input int64 = 12345
	var encoded = Encode(input)
	var decoded = Decode(encoded)
	if decoded != input {
		t.Errorf("Encode(%d) = %s, Decode(%s) = %d", input, encoded, encoded, decoded)
	}
}

func TestDecode(t *testing.T) {
	var input = "4T7HwLG6"
	var decoded = Decode(input)
	var encoded = Encode(decoded)
	if encoded != input {
		t.Errorf("Decode(%s) = %d, Encode(%d) = %s", input, decoded, decoded, encoded)
	}
}
