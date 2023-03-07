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
	"web/internal/apperror"
	"web/internal/entity"
	"web/internal/handlers"
	"web/internal/user/usecase"
)

type handler struct {
	service *usecase.Service
}

const (
	startPage   = "/"
	login       = "/login"
	signup      = "/signup"
	dashboard   = "/dashboard"
	users       = "/users"
	usersId     = "/users/:userId"
	usersCreate = "/users/create"
	mainPage    = "/general"
)

func NewHandler(service *usecase.Service) handlers.Handler {
	return &handler{service: service}
}

func (h *handler) Register(router *httprouter.Router) {

	router.ServeFiles("/public/*filepath", http.Dir("public"))

	router.GET(startPage, h.StartPage)
	router.POST(login, h.Login)
	router.POST(signup, h.SignUp)
	router.GET(dashboard, apperror.AuthMiddleware(h.AccountPage))
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

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == "POST" {
		login := r.FormValue("login")
		password := r.FormValue("password")

		user := &entity.User{
			Login:    login,
			Password: password,
		}

		err := h.service.SignUp(r.Context(), user)
		if err != nil {
			log.Fatalf("failed to get method SignUp")
		}
		//allBytes, err := json.MarshalIndent(user, "", "")
		//if err != nil {
		//	fmt.Printf("Error : %v", err)
		//}
		//w.WriteHeader(201)
		//fmt.Fprintf(w, "Create user: %s \n", allBytes)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

}
func (h *handler) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == "POST" {

		login := r.FormValue("login")
		password := r.FormValue("password")

		users, err := h.service.LogIn(r.Context(), login, password)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		userToken := users.Token

		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    userToken,
			Expires:  time.Now().Add(time.Minute * 60),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		//path := filepath.Join("public", "index.html")
		//tmpl, err := template.ParseFiles(path)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//	return
		//}
		//
		//err = tmpl.ExecuteTemplate(w, "index", nil)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//	return
		//	h.StartPage(w, r, p)
	}
}

func (h *handler) UpdateDataUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	w.Write([]byte("update data user"))
}
func (h *handler) PartialUpdateDataUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	w.Write([]byte("partial update data user"))
}
func (h *handler) StartPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	w.Write([]byte("delete user"))
}
