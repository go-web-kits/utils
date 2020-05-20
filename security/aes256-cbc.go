package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func NotEncoded(s string) bool {
	if s == "" {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(s)
	return err != nil
}

func Encrypt(data, key string) (string, error) {
	if data == "" {
		return "", nil
	}

	encrypted, err := AESxCBC_Encrypt([]byte(data), []byte(key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func Decrypt(data, key string) (string, error) {
	if data == "" {
		return "", nil
	}

	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	decrypted, err := AESxCBC_Decrypt(decoded, []byte(key))
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

// ======

// TODO refactor

// 使用 PKCS7 进行填充，IOS 也是7
func paddingPKCS7(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 解填充
func unpaddingPKCS7(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 填充秘钥 key 16, 24, 32 位分别对应 AES-128, AES-192, AES-256
func AESxCBC_Encrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 填充原文
	blockSize := block.BlockSize()
	data = paddingPKCS7(data, blockSize)
	cipherText := make([]byte, len(data))
	iv := make([]byte, blockSize)

	// block 大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, data)

	return cipherText, nil
}

func AESxCBC_Decrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()

	if len(data) < blockSize {
		panic("ciphertext too short")
	}
	iv := make([]byte, blockSize)

	// CBC mode always works in whole blocks.
	if len(data)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(data, data)
	return unpaddingPKCS7(data), nil
}
