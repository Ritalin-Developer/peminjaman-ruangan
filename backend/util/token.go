package util

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/config"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

// DecodePrivateECDSA is a function to decode private ECDSA
// func DecodePrivateECDSA(pemEncoded string) (privateKey *ecdsa.PrivateKey, err error) {
// 	block, rest := pem.Decode([]byte(pemEncoded))
// 	if len(rest) != 0 && block == nil {
// 		err = fmt.Errorf("invalid private_key format")
// 		log.Error(err)
// 		return
// 	}
// 	x509Encoded := block.Bytes
// 	privateKey, _ = x509.ParseECPrivateKey(x509Encoded)
// 	return
// }

// GenerateToken is a utility function to generate JWT token
func GenerateToken(tokenData jwt.MapClaims) (tokenString string, err error) {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}
	expireAt := time.Now().Add(time.Minute * time.Duration(config.TokenLifetimeMin)).Unix()
	tokenData["expire_at"] = expireAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData)
	// claims := token.Claims.(jwt.MapClaims)

	// claims["authorized"] = true
	// claims["username"] = username
	// claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err = token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (valid bool, claims jwt.MapClaims, err error) {
	// config, err := config.LoadConfig(".")
	// if err != nil {
	// 	log.Fatalf("can't load environment app.env: %v", err)
	// }

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if err != nil {
			log.Error(err)
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		token.Valid = true
		return token, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var exp int64
		switch e := claims["expire_at"].(type) {
		case string:
			conv, err := strconv.Atoi(e)
			if err != nil {
				log.Error(err)
				break
			}
			exp = int64(conv)
		case float64:
			exp = int64(e)
		}
		if exp < time.Now().Unix() {
			return false, claims, errors.New("expired token")
		}
		return true, claims, nil
	}
	return
}
