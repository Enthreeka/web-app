package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"web/internal/account/usecase"
	"web/internal/apperror"
	"web/internal/entity"
	"web/internal/handlers"
)

type handler struct {
	service *usecase.Service
}

func NewAccountHandler(service *usecase.Service) handlers.Handler {
	return &handler{
		service: service,
	}
}

type Task struct {
	Name        string
	Description string
}

func (h *handler) Register(router *httprouter.Router) {
	log.Println("Registering account routes...")

	router.POST("/dashboard/add", h.AddTask)
	router.GET("/dashboard", apperror.AuthMiddleware(h.GetTask))
}

func (h *handler) GetTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	log.Printf("Handling GetTask request with parameters: %v", p)
	path := filepath.Join("public", "index2.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//get username from cookie for transfers to html
	cookie, err := r.Cookie("username")
	if err != nil {
		log.Fatalf("failed to get cookie")
		return
	}
	cookieUsername := cookie.Value

	q := r.URL.Query()
	userID, _ := strconv.Atoi(q.Get("id"))

	name, description, err := h.service.GetTask(r.Context(), userID)
	if name != nil || description != nil {
		if err != nil {
			log.Printf("failed with get task in handler %v", err)
			return
		}

		type DataUser struct {
			Tasks    []Task
			Username string
		}

		tasks := []Task{}
		for i, name := range name {
			tasks = append(tasks, Task{Name: name, Description: description[i]})
		}

		data := DataUser{
			Tasks:    tasks,
			Username: cookieUsername,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Fatalf("COULD NOT EXECUTE %v", err)
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		//if new user or user without fields,then he got this clear page
		type DataUser struct {
			Tasks    []Task
			Username string
		}
		data := DataUser{
			Username: cookieUsername,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Fatalf("COULD NOT EXECUTE %v", err)
			http.Error(w, err.Error(), 400)
			return
		}
	}

}

func (h *handler) AddTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("Handling UpdateDescriptionTask request with parameters: %v", p)

	descriptionName := r.FormValue("descriptionName")
	description := r.FormValue("description")

	//q := r.URL.Query()
	//userID, err := strconv.Atoi(q.Get("id"))
	//if err != nil || userID <= 1 {
	//	log.Printf("failed to get id from url in account/handler %v", err)
	//	http.NotFound(w, r)
	//	return
	//}
	//fmt.Println(userID)

	task := &entity.Task{
		AccountId:       2,
		NameTask:        descriptionName,
		DescriptionTask: description,
	}

	err := h.service.CreateTask(r.Context(), task)
	if err != nil {
		fmt.Printf("failed to add taks %v", err)
		return
	}

	q := url.Values{}
	q.Add("id", strconv.Itoa(2))
	url := fmt.Sprintf("/dashboard?%s", q.Encode())

	http.Redirect(w, r, url, http.StatusSeeOther)
}
