package http

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
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

	router.GET(startPage, h.StartPage)
	router.POST(login, h.Login)
	router.POST(signup, h.SignUp)
	//	router.GET(dashboard, apperror.AuthMiddleware(h.AccountPage))
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == "POST" {
		login := r.FormValue("login")
		if !validation.IsValidationLogin(login) {
			fmt.Println("login did not meet the requirements")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		password := r.FormValue("password")
		if !validation.IsValidationPassword(password) {
			fmt.Println("password did not meet the requirements")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		user := &entity.User{
			Login:    login,
			Password: password,
		}

		dataUser, err := h.service.SignUp(r.Context(), user)
		if err != nil {
			log.Printf("failed to get method SignUp error:%v", err)
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "username",
			Value: dataUser.Login,
		})

		q := url.Values{}
		q.Add("id", strconv.Itoa(dataUser.Id))
		url := fmt.Sprintf("/dashboard?%s", q.Encode())

		fmt.Println("test")
		http.Redirect(w, r, url, http.StatusSeeOther)
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

		//q := r.URL.Query()
		//q.Set("user_id", strconv.Itoa(users.Id))
		//query := q.Encode()
		//userID := p.ByName("user_id")
		//
		//fmt.Println("Encode - ", query)
		//fmt.Println("method Login - ", userID)

		h.user.Login = login
		h.user.Id = users.Id

		userToken := users.Token

		//set in cookie jwt token for 24 hours
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    userToken,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

		//set in cookie username/login
		http.SetCookie(w, &http.Cookie{
			Name:  "username",
			Value: users.Login,
		})

		//Set the user_id in url
		q := url.Values{}
		q.Add("id", strconv.Itoa(users.Id))
		url := fmt.Sprintf("/dashboard?%s", q.Encode())

		http.Redirect(w, r, url, http.StatusSeeOther)
	}
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

	data := AccountPageData{
		User: h.user,
	}

	fmt.Println(data.User.Id)
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
