package encoder

import (
    "testing"
)

func TestBase64ToInt(t *testing.T) {
    type args struct {
        base64 string
    }
    var tests []struct {
        name string
        args args
        want int
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Base64ToInt(tt.args.base64); got != tt.want {
                t.Errorf("Base64ToInt() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBase64ToUInt(t *testing.T) {
    type args struct {
        base64 string
    }
    tests := []struct {
        name string
        args args
        want uint
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Base64ToUInt(tt.args.base64); got != tt.want {
                t.Errorf("Base64ToUInt() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestBase64ToUInt64(t *testing.T) {
    type args struct {
        base64 string
    }
    tests := []struct {
        name string
        args args
        want uint64
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Base64ToUInt64(tt.args.base64); got != tt.want {
                t.Errorf("Base64ToUInt64() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestIntToBase64(t *testing.T) {
    type args struct {
        number int
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := IntToBase64(tt.args.number); got != tt.want {
                t.Errorf("IntToBase64() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestUInt64ToBase64(t *testing.T) {
    type args struct {
        number uint64
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := UInt64ToBase64(tt.args.number); got != tt.want {
                t.Errorf("UInt64ToBase64() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestUIntToBase64(t *testing.T) {
    type args struct {
        number uint
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := UIntToBase64(tt.args.number); got != tt.want {
                t.Errorf("UIntToBase64() = %v, want %v", got, tt.want)
            }
        })
    }
}
