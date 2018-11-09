package Services

import (
	"fmt"
	"testing"
)

func TestEncryption(t *testing.T) {
	str := "asdasdasd"
	Enc_str := Encryption(str)
	Dec_str := Decryption(Enc_str)

	fmt.Println(Dec_str)
	fmt.Println(Enc_str)
}
