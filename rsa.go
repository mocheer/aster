package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 解析公匙
func RSA_PublicKeyFromBytes(pubByte []byte) (*rsa.PublicKey, error) {
	// pem解码
	b, _ := pem.Decode(pubByte)
	if b == nil {
		return nil, errors.New("error public key")
	}
	// der解码，最终返回一个公匙对象
	// pubKey, err := x509.ParsePKCS1PublicKey(b.Bytes)
	pubKey, err := x509.ParsePKIXPublicKey(b.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey.(*rsa.PublicKey), nil
}

//  rsa公匙加密
func RSA_Encrypt(src []byte, publickey *rsa.PublicKey) ([]byte, error) {
	// 使用公匙加密数据，需要一个随机数生成器和公匙和需要加密的数据
	data, err := rsa.EncryptPKCS1v15(rand.Reader, publickey, src)
	if err != nil {
		return nil, err
	}
	return data, nil
}
