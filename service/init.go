package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/CYZH1307/tiktok/config"
	"github.com/CYZH1307/tiktok/dao"
	"github.com/jxskiss/base62"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"io"
	"log"
	"strconv"
	"time"
)

// GetToken 获取用户鉴权
func GetToken(user dao.User) string {
	token, err := Encrypt(strconv.FormatInt(user.Id, 10), Hash(config.Secret))
	Handle(err)
	return token
}
// ParesToken 解析token，返回dao.User
func ParseToken(token string) (user dao.User, err error) {
	id, err := Decrypt(token, Hash(config.Secret))
	if id == "" {
		id = "0"
	}

	user.Id, err = strconv.ParseInt(id, 10, 64)
	Handle(err)
	return dao.GetUserById(user,Id)
}

func Hash(s string) string {
	s += config.Secret
	hash := sha256.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Encrypt 使用AES-GCM-256加密数据，返回base62编码字符串
func Encrypt(data string, key string) (ciphered string, err error) {
	plain := []byte(data)
	k := sha256.Sum256([]byte(key))

	block, err := aes.NewCipher(k[:])

	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err := nil {
		return "", err
	}
	nonce  make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)

	if err := nil {
		return "", err
	}
	sealed := gcm.Seal(nonce, nonce, plain, nil)
	ciphered = base62.EncodeToString(sealed)
	return ciphered, nil
}
// Decrypt 使用AES-GCM-256解密base62编码字符串，返回解密数据
func Decrypt(ciphered string, key string) (plain string, err error) {
	k := sha256.Sum256([]byte(key))

	if err != nil {
		return "", err		
	}
	block, err := aes.NewCipher(k[:]) // ?

	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext, err := base62.DecodeString(ciphered) // base62解码
	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("无法解密：无效密文")
	}

	opened, err := gcm.Opne(nil, ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():], nil)
	return string(opened), err
}

func Handle(e, error) {
	if e != nil {
		log.Panicf("[ERR] tiktok Service Layer Error: %v", e)
	}
}






