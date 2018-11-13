package Services

import (
	"github.com/kylesliu/gin-demo/App/Extensions/Crypto"
	"github.com/kylesliu/gin-demo/Bootstrap/config"
)

//	加密
func Encryption(str string) string {
	key := config.AppConfig.EncryptKey
	res := Crypto.EncryptDES_ECB(str, key)
	return res
}

//	解密
func Decryption(str string) string {
	key := config.AppConfig.EncryptKey
	res := Crypto.DecryptDES_ECB(str, key)
	return res
}
