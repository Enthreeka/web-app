package apperror

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		if !isAuthorized(r) {
			log.Printf("failed in Authorized")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r, p)
	}
}

func isAuthorized(r *http.Request) bool {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		log.Printf("failed to get jwt index %v", err)
		return false
	}
	token := cookie.Value

	if !isValidToken(token) {
		log.Printf("failed to valid jwt token")
		return false
	}
	return true
}

func isValidToken(token string) bool {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-token-gen"), nil
	})
	if err != nil {
		log.Printf("failed to pare token %v", err)
		return false
	}
	if parsedToken.Valid {
		expTime := parsedToken.Claims.(jwt.MapClaims)["exp"].(float64)
		if int64(expTime) < time.Now().Unix() {
			log.Printf("jwt token has expired")
			return false
		}
		return true
	} else {
		log.Printf("jwt token not valid")
		return false
	}
}
