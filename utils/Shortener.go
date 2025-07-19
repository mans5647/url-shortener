package utils

import (
	"crypto/md5"
	"strconv"
	"encoding/hex"
	"github.com/deatil/go-encoding/encoding"
)


func getBytesFromHex(hexStr string, howMuch int) []byte {

	bytes := make([]byte, 10)

	for i := 0; i < howMuch * 2; i++ {
		bytes = append(bytes, hexStr[i])
	}
	return bytes
}

func convertHexBytesToNumber(hexBytes []byte) int {

	num, _ := strconv.ParseInt(string(hexBytes), 16, 32)
	return int(num)
}

/*
	Generates short code for number using Base62 encoding
*/
func GenerateShortCode(nextId int, url string) string {
	
	md := md5.New()

	data := md.Sum([]byte(url))

	hashString := hex.EncodeToString(data)

	finalValue := nextId + convertHexBytesToNumber(getBytesFromHex(hashString, 6))

	return encoding.FromString(strconv.Itoa(finalValue)).Base62Encode().ToString()

}