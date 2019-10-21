package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/iceCream/src/constants"
	"github.com/iceCream/src/web/controller"
)

//SetupRoutes sets up all incoming request urls for application
var SetupRoutes = func() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc(constants.AddIC, ValidateMiddleware(controller.InsertIceCream)).Methods("Post")
	//router.Handle("constants.AddICs", controller.InsertIceCreams).Methods("Post")
	router.HandleFunc(constants.GetAllIc, ValidateMiddleware(controller.GetAllIceCreams)).Methods("Get")
	router.HandleFunc(constants.GetICByProductID, ValidateMiddleware(controller.GetIceCreamByProductID)).Methods("Get")
	router.HandleFunc(constants.Token, controller.GenerateToken).Methods("Post")

	return router

}

//ValidateMiddleware intercepts each request
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(error.Error())
					return
				}
				if token.Valid {
					//context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode("Invalid authorization token")
				}
			}
		} else {
			json.NewEncoder(w).Encode("An authorization header is required")
		}
	})
}
