package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
	"web/internal/entity"
	"web/internal/handlers"
	"web/internal/user/usecase"
	"web/internal/validation"
)

type handler struct {
	service *usecase.Service
	user    *entity.User
}

type AccountPageData struct {
	User *entity.User
}

const (
	startPage = "/"
	login     = "/login"
	signup    = "/signup"
	dashboard = "/dashboard"
)

func NewHandler(service *usecase.Service, user *entity.User) handlers.Handler {
	return &handler{
		service: service,
		user:    user,
	}
}

func (h *handler) Register(router *httprouter.Router) {

	router.ServeFiles("/public/*filepath", http.Dir("public"))

	router.GET(startPage, h.StartPageHandler)
	router.POST(login, h.LoginPageHandler)
	router.POST(signup, h.SignUpPageHandler)
}

func (h *handler) SignUpPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == "POST" {
		login := r.FormValue("login")
		if !validation.IsValidationLogin(login) {
			fmt.Println("login did not meet the requirements")
			http.Redirect(w, r, startPage, http.StatusSeeOther)
			return
		}
		password := r.FormValue("password")
		if !validation.IsValidationPassword(password) {
			fmt.Println("password did not meet the requirements")
			http.Redirect(w, r, startPage, http.StatusSeeOther)
			return
		}

		user := &entity.User{
			Login:    login,
			Password: password,
		}

		dataUser, err := h.service.SignUp(r.Context(), user)
		if err != nil {
			log.Printf("failed to get method SignUpPageHandler error:%v", err)
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "username",
			Value: dataUser.Login,
		})

		http.SetCookie(w, &http.Cookie{
			Name:  "id",
			Value: dataUser.Id,
		})

		http.Redirect(w, r, dashboard, http.StatusSeeOther)
	}

}
func (h *handler) LoginPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == "POST" {
		login := r.FormValue("login")
		password := r.FormValue("password")

		users, err := h.service.LogIn(r.Context(), login, password)
		if err != nil {
			http.Redirect(w, r, startPage, http.StatusSeeOther)
		}

		h.user.Login = login
		h.user.Id = users.Id

		userToken := users.Token

		//set in cookie jwt token for 24 hours
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    userToken,
			Expires:  time.Now().Add(time.Hour * 1),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

		//set in cookie username/login
		http.SetCookie(w, &http.Cookie{
			Name:  "username",
			Value: users.Login,
		})

		http.SetCookie(w, &http.Cookie{
			Name:  "id",
			Value: users.Id,
		})

		http.Redirect(w, r, dashboard, http.StatusSeeOther)
	}
}
func (h *handler) StartPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := filepath.Join("public", "index.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}
