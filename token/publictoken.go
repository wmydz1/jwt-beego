package token

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// location of the files used for signing and verification
const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// keys are held in global variables
// i havn't seen a memory corruption/info leakage in go yet
// but maybe it's a better idea, just to store the public key in ram?
// and load the signKey on every signing request? depends on  your usage i guess
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// read the key files before starting http handlers
func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}
func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewToken() string {
	// create a signer for rsa 256
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["username"] = "samchen";
	claims["password"] = "chen168";
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(signKey)
	if err == nil {
		return tokenString
	}
	return "something error"
}
func ValidateToken(tokenString string) bool {
	// validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	// branch out into the possible error from signing
	switch err.(type) {

	case nil: // no error

		if !token.Valid { // but may still be invalid
			log.Println("WHAT? Invalid Token? F*** off!")
			return false
		}
	case *jwt.ValidationError: // something was wrong during the validation
		vErr := err.(*jwt.ValidationError)
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			log.Println("Token Expired, get a new one.")
			return false
		default:
			log.Printf("ValidationError error: %+v\n", vErr.Errors)
			return false
		}
	default: // something else went wrong
		log.Printf("Token parse error: %v\n", err)
		return false
	}
	if err == nil && token.Valid {
		return true
	}
	return false

}
