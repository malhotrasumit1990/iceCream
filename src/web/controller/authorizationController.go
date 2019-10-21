package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	Token string `json:"token"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateJWT(w http.ResponseWriter, req *http.Request) JwtToken {
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	return JwtToken{Token: tokenString}
}

/*
func GenerateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Ben & Jerry"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	//should be read from env Variable
	tokenString, err := token.SignedString("supersecretkey")
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
*/
var GenerateToken = func(w http.ResponseWriter, r *http.Request) {

	tokenStruct := GenerateJWT(w, r)

	json.NewEncoder(w).Encode(tokenStruct)

}
