package user

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	"web/internal/entity"
	"web/internal/handlers"
)

type handler struct {
	service *Service
}

const (
	users       = "/users"
	usersId     = "/users/:userId"
	usersCreate = "/users/create"
	startPage   = "/"
	mainPage    = "/general"
	login       = "/login"
)

func NewHandler(service *Service) handlers.Handler {
	return &handler{service: service}
}

func (h *handler) Register(router *httprouter.Router) {

	router.ServeFiles("/public/*filepath", http.Dir("public"))

	//router.GET(startPage, h.StartPage)
	router.GET(mainPage, h.MainPage)
	router.GET(users, h.GetUsers)
	router.POST(usersCreate, h.CreateUser)
	router.GET(login, h.StartPage)
	//router.PUT(usersId, h.UpdateDataUser)
	//router.PATCH(usersId, h.PartialUpdateDataUser)
	//router.DELETE(usersId, h.DeleteUser)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	login := r.FormValue("login")
	password := r.FormValue("password")

	user := &entity.User{
		Login:    login,
		Password: password,
	}

	err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	allBytes, err := json.MarshalIndent(user, "", "")
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, "Create user: %s \n", allBytes)

}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//userId, err := strconv.Atoi(strings.TrimLeft(p.ByName("userId"), ":")) <---Get user by id
	//if err != nil {
	//	fmt.Println(err)
	//}
	path := filepath.Join("public", "index.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	users, err := h.service.GetOne(r.Context(), login, password)
	if err != nil {
		w.WriteHeader(404) //TODO set new status.code
		return
	}
	data := struct {
		Login    string
		Password string
	}{
		Login:    users.Login,
		Password: users.Password,
	}
	//
	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}
func (h *handler) UpdateDataUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("update data user"))
}
func (h *handler) PartialUpdateDataUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("partial update data user"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("delete user"))
}

func (h *handler) StartPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")

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
	w.Header().Set("Content-Type", "text/html")

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
