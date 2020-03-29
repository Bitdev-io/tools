package encryptor

import (
    "crypto"
    "reflect"
    "testing"
)

func TestEncryptor_Decrypt(t *testing.T) {
    type fields struct {
        Encryption crypto.Hash
    }
    type args struct {
        key []byte
        b64 string
    }
    var tests []struct {
        name   string
        fields fields
        args   args
        want   string
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            encryptor := &Encryptor{
                Encryption: tt.fields.Encryption,
            }
            if got := encryptor.Decrypt(tt.args.key, tt.args.b64); got != tt.want {
                t.Errorf("Decrypt() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestEncryptor_Encrypt(t *testing.T) {
    type fields struct {
        Encryption crypto.Hash
    }
    type args struct {
        key  []byte
        text []byte
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            encryptor := &Encryptor{
                Encryption: tt.fields.Encryption,
            }
            if got := encryptor.Encrypt(tt.args.key, tt.args.text); got != tt.want {
                t.Errorf("Encrypt() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewEncryptor(t *testing.T) {
    tests := []struct {
        name string
        want *Encryptor
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewEncryptor(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewEncryptor() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewEncryptorWithDefault(t *testing.T) {
    type args struct {
        defaultEncryption *Encryptor
    }
    tests := []struct {
        name string
        args args
        want *Encryptor
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewEncryptorWithDefault(tt.args.defaultEncryption); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewEncryptorWithDefault() = %v, want %v", got, tt.want)
            }
        })
    }
}
