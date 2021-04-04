package login

import (
	"crypto/rsa"
	"flaber-auth/internal/env"
	"flaber-auth/internal/logger"
	"flaber-auth/internal/models"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type UserToken models.User

var (
	signKey *rsa.PrivateKey
	privateKey string
)

// JWT personzalizado
type jwtCustomClaims struct {
	User      *models.User `json:"user"`
	IPAddress string      `json:"ip_address"`
	jwt.StandardClaims
}

// init lee los archivos de firma y validación RSA
func init() {
	c := env.NewConfiguration()
	privateKey = c.App.RSAPrivateKey
	signBytes, err := ioutil.ReadFile(privateKey)
	if err != nil {
		logger.Error.Printf("leyendo el archivo privado de firma: %s", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logger.Error.Printf("realizando el parse en authentication RSA private: %s", err)
	}
}

// Genera el token
func GenerateJWT(u *models.User) (string, int, error) {
	c := &jwtCustomClaims{
		User:      u,
		IPAddress: u.RealIP,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1200).Unix(),
			Issuer:    "Fluber",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	token, err := t.SignedString(signKey)
	if err != nil {
		logger.Error.Printf("firmando el token: %v", err)
		return "", 70, err
	}

	return token, 29, nil
}