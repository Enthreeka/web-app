package http

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
	"web/internal/entity"
	"web/internal/handlers"
	"web/internal/user/usecase"
)

type handler struct {
	service *usecase.Service
}

const (
	users       = "/users"
	usersId     = "/users/:userId"
	usersCreate = "/users/create"
	startPage   = "/"
	mainPage    = "/general"
	login       = "/login"
	dashboard   = "/dashboard"
)

func NewHandler(service *usecase.Service) handlers.Handler {
	return &handler{service: service}
}

func (h *handler) Register(router *httprouter.Router) {

	router.ServeFiles("/public/*filepath", http.Dir("public"))

	router.GET(startPage, h.StartPage)
	router.POST(login, h.Login)
	router.GET(dashboard, h.MainPage)
	//	router.GET(login, h.StartPage)
	//router.PUT(usersId, h.UpdateDataUser)
	//router.PATCH(usersId, h.PartialUpdateDataUser)
	//router.DELETE(usersId, h.DeleteUser)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	all, err := h.service.GetAll(r.Context())

	if err != nil {
		w.WriteHeader(400)
		return
	}

	allBytes, err := json.MarshalIndent(all, "", "")
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	w.WriteHeader(200)
	w.Write(allBytes)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	login := r.FormValue("login")
	password := r.FormValue("password")

	user := &entity.User{
		Login:    login,
		Password: password,
	}

	//err := h.service.CreateUser(r.Context(), user)
	//if err != nil {
	//	w.WriteHeader(404)
	//	return
	//}

	allBytes, err := json.MarshalIndent(user, "", "")
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, "Create user: %s \n", allBytes)

}

func (h *handler) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {

		login := r.FormValue("login")
		password := r.FormValue("password")

		users, err := h.service.LogIn(r.Context(), login, password)
		if err != nil {
			http.Error(w, "invalid login or password", http.StatusUnauthorized)
			w.WriteHeader(401)
			return
		}
		userToken := users.Token

		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    userToken,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		path := filepath.Join("public", "index.html")
		tmpl, err := template.ParseFiles(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = tmpl.ExecuteTemplate(w, "index", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
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
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	w.Write([]byte("delete user"))
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

func (h *handler) MainPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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
