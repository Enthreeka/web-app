package http

import (
	"encoding/json"
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

	//router.GET(dashboard, apperror.AuthMiddleware(h.AccountPage))
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
func (h *handler) AccountPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := filepath.Join("public", "index2.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	data := AccountPageData{
		User: h.user,
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: data.User.Login,
	})

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	http.Redirect(w, r, "/dashboard/add", http.StatusSeeOther)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	all, err := h.service.GetAll(r.Context())

	if err != nil {
		w.WriteHeader(400)
		return nil
	}

	allBytes, err := json.MarshalIndent(all, "", "")
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	w.WriteHeader(200)
	w.Write(allBytes)
	return nil
}
