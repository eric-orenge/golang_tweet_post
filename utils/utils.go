package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

//RespondErr response handling
func RespondErr(w http.ResponseWriter, r *http.Request, message string) error {
	resp := map[string]interface{}{
		"success": 0,
		"message": message,
	}
	return json.NewEncoder(w).Encode(resp)
}
func RespondSuccess(w http.ResponseWriter, r *http.Request, data interface{}) error {
	resp := map[string]interface{}{
		"success": 1,
		"message": data,
	}
	return json.NewEncoder(w).Encode(resp)
}
func GenerateKey() (int, error) {
	key := make([]byte, 64)

	oauth_nonce, err := rand.Read(key)
	if err != nil {
		return 0, err
		log.Fatal("Err generating key")
	}
	return oauth_nonce, nil
}

//password encryption and decryption

func Encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}
func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
