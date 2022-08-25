package libs

import (
	"fmt"

	"crypto/md5"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

//fungsi bcrypt
func HashBcryptPassword(password string) (string, error) { //algoritma bcrypt berjalan lebih lambat daripada md5
	//bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CompareBcrypt(hash1, hash2 []byte) bool { //dalam bcrypt compare hnay bisa antara hash dengna plaintext
	fmt.Println("isi hash 1", hash1)
	fmt.Println("isi hash 2", hash2)
	var flag bool
	err := bcrypt.CompareHashAndPassword(hash1, hash2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success sama")
		flag = true
	}
	return flag
}

//fungsi MD5
func HashMD5Password(password string) string {
	hashString := []byte(password)
	hash_value := md5.Sum(hashString)
	return hex.EncodeToString(hash_value[:])
}

//fungsi base 64 encode
func Base64encode(plaintext string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(plaintext))
	return encoded
}

//fungsi base 64 decode
func Base64decode(encodetext string) string {
	decoded, err := base64.StdEncoding.DecodeString(encodetext)
	if err != nil {
		fmt.Println("decode error:", err)

	}
	fmt.Println("hasil decode", string(decoded))
	return string(decoded)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
