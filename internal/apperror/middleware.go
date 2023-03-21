package apperror

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

//type appHandl func(w http.ResponseWriter, r *http.Request, p httprouter.Params)

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

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				} else if errors.Is(err, NoAuthErr) {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(ErrNotFound.Marshal())
					return
				}

				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(ErrNotFound.Marshal())
			}

			w.WriteHeader(http.StatusTeapot) //418
			w.Write(systemError(err).Marshal())
		}
	}
}
