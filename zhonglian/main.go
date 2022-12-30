package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

type AesCryptor struct {
	key []byte
	iv  []byte
}

func (a *AesCryptor) Encrypt(data []byte) ([]byte, error) {
	aesBlockEncrypter, err := aes.NewCipher(a.key)
	content := PKCS5Padding(data, aesBlockEncrypter.BlockSize())
	encrypted := make([]byte, len(content))
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, a.iv)
	aesEncrypter.CryptBlocks(encrypted, content)
	return encrypted, nil
}

//解密数据
func (a *AesCryptor) Decrypt(src []byte) (data []byte, err error) {
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher(a.key)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesDecrypter := cipher.NewCBCDecrypter(aesBlockDecrypter, a.iv)
	aesDecrypter.CryptBlocks(decrypted, src)
	return PKCS5Trimming(decrypted), nil
}

/**
PKCS5包装
*/
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

/*
解包装
*/
func PKCS5Trimming(encrypt []byte) []byte {
	return encrypt[:11]
}

func main() {
	//key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
	token := "Kov2BA2cii+hCK4Yt3G7PIfhxzlZ6s5FsU67wp1EMWY="
	//token := "7xu/86C7OK2gMZdYVgf6hlAY7dKCt8f7xabTEELQNag="
	//token := "qoRlvaLf9zMvR7zaL5N7jqKC9nvehnEaVEpa6N3pE1s="
	//token := "1QVRVBQ+bZQaehaYfiJUTQ=="
	c := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//d := []byte{1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6}
	e := []byte("1234567890123456")
	//f := []byte("0000000000000000")
	//phoneNum := []byte("13580000000")
	//phoneNum := []byte("18860230359")

	a := AesCryptor{
		//key: d,
		iv:  c,
		key: e,
		//iv:  f,
	}

	//xpass, err := a.Encrypt(phoneNum)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//pass64 := base64.StdEncoding.EncodeToString(xpass)
	//fmt.Printf("加密后:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	tpass, err := a.Decrypt(bytesPass)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Printf("解密后:%v\n", string(tpass))
}
