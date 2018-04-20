package validate

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(){
	tokenString:="1"
	hmacSampleSecret :=[]byte("logoocc.com")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})
	if err ==nil{
		fmt.Println(token.Valid)
	}
	fmt.Println("validate error: ",err)
}
