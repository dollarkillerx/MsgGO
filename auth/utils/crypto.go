package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

func Md5Encode(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}


/**
Rsa256加密
param: origData原始数据
param: publicKey 公钥
 */
func Rsa256Encrypt(origData,publicKey[]byte) ([]byte,error) {
	p, _ := pem.Decode(publicKey)
	if p == nil {
		return nil,errors.New("public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil{
		return nil,err
	}

	key := pub.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader,key,origData)
}

/**
Rsa256解密
param: ciphertext 密文
param: privateKey 私钥
 */
func Rsa256Decrypt(ciphertext,privateKey []byte) ([]byte,error) {
	p, _ := pem.Decode(privateKey)
	if p == nil {
		return nil,errors.New("private key error")
	}
	key, e := x509.ParsePKCS1PrivateKey(p.Bytes)
	if e != nil {
		return nil,e
	}
	return rsa.DecryptPKCS1v15(rand.Reader,key,ciphertext)
}