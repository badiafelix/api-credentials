package libs

import (
	"fmt"

	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

//fungsi bcrypt
func HashBcryptPassword(password string) (string, error) { //algoritma bcrypt berjalan lebih lambat daripada md5
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//fungsi MD5
func HashMD5Password(password string) string {
	hashString := []byte(password)

	hash_value := md5.Sum(hashString)
	fmt.Printf("hasil md5: %x", md5.Sum(hashString))
	//bytes := md5.Sum([]byte(password))
	return hex.EncodeToString(hash_value[:])
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
