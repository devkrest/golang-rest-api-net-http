package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

var (
	algorithmKey = []byte("1234567890123456") // 16 / 24 / 32 bytes
	initVector   = []byte("1234567890123456") // 16 bytes
)

func DecryptKey(key string) (string, error) {

	data, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(algorithmKey)
	if err != nil {
		return "", err
	}

	if len(data)%aes.BlockSize != 0 {
		return "", errors.New("invalid block size")
	}

	mode := cipher.NewCBCDecrypter(block, initVector)
	mode.CryptBlocks(data, data)

	// remove padding
	length := len(data)
	unpad := int(data[length-1])
	return string(data[:(length - unpad)]), nil
}
