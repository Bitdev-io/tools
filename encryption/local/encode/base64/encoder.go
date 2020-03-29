package encoder

import "strings"

var b64 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
var base = len(b64)

func IntToBase64(number int) string {
	str := ""
	for number > 0 {
		r := number % base
		number -= r
		number /= base
		str = string(b64[r]) + str
	}
	return str
}

func Base64ToInt(base64 string) int {
	number := 0
	sliced := strings.Split(base64, "")
	for _, v := range sliced {
		r := strings.Index(b64, v)
		number *= base
		number += r
	}
	return number
}

func UInt64ToBase64(number uint64) string {
	str := ""
	for number > 0 {
		r := number % uint64(base)
		number -= r
		number /= uint64(base)
		str = string(b64[r]) + str
	}
	return str
}

func Base64ToUInt64(base64 string) uint64 {
	number := uint64(0)
	sliced := strings.Split(base64, "")
	for _, v := range sliced {
		r := strings.Index(b64, v)
		number *= uint64(base)
		number += uint64(r)
	}
	return number
}

func UIntToBase64(number uint) string {
	str := ""
	for number > 0 {
		r := number % uint(base)
		number -= r
		number /= uint(base)
		str = string(b64[r]) + str
	}
	return str
}

func Base64ToUInt(base64 string) uint {
	number := uint(0)
	sliced := strings.Split(base64, "")
	for _, v := range sliced {
		r := strings.Index(b64, v)
		number *= uint(base)
		number += uint(r)
	}
	return number
}
