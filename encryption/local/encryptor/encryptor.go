package encryptor

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
)


type Encryptor struct {
	Encryption crypto.Hash
}


var hash = crypto.SHA512
var DefaultEncryption = &Encryptor{
	Encryption: hash,
	}

func init() {
	fmt.Println("loaded")
}

// NewEncryptor  get new encryptor instance
func NewEncryptor() *Encryptor {
	return NewEncryptorWithDefault(DefaultEncryption)
}

// NewEncryptorWithDefault  Sets up a default state encryptor in version v0.0.1 encryptor is .ini
func NewEncryptorWithDefault(defaultEncryption *Encryptor) *Encryptor {
	return &Encryptor{
		Encryption: DefaultEncryption.Encryption,
	}
}


func (encryptor *Encryptor) Encrypt(key, text []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(crand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return encodeBase64(ciphertext)
}


func (encryptor *Encryptor) Decrypt(key []byte, b64 string) string {
	text := encoder.Base64ToInt(b64)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(text)
}
