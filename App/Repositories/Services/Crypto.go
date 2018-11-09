package Services

import "github.com/kylesliu/gin-demo/App/Extensions/Crypto"

//	加密
func Encryption(str string) string {
	key := "o6uoO27*"
	res := Crypto.EncryptDES_ECB(str, key)
	return res
}

//	解密
func Decryption(str string) string {
	key := "o6uoO27*"
	res := Crypto.DecryptDES_ECB(str, key)
	return res
}
