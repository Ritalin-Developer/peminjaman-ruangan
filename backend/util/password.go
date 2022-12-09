package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePassword := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
