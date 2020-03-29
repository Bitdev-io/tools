package encryptor

import "fmt"

type Encryptor struct {
	Encryption string

}

var DefaultEncryption = &Encryptor{
	Encryption: "ini",
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
