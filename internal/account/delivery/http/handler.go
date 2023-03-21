package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
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
	Id          string
}

func (h *handler) Register(router *httprouter.Router) {
	log.Println("Registering account routes...")

	router.POST("/dashboard/add", h.AddTask)
	router.GET("/dashboard", apperror.AuthMiddleware(h.GetTask))
	router.DELETE("/dashboard/delete", h.DeleteTask)
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
	cookieUN, err := r.Cookie("username")
	if err != nil {
		log.Fatalf("failed to get cookie %v", err)
		return
	}
	cookieUserID, err := r.Cookie("id")
	if err != nil {
		log.Fatalf("failed to get cookie %v", err)
		return
	}
	cookieID := cookieUserID.Value
	cookieUsername := cookieUN.Value

	//q := r.URL.Query()
	//userID := q.Get("id")
	//
	//fmt.Println(userID)

	id, name, description, err := h.service.GetTask(r.Context(), cookieID)
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
			tasks = append(tasks, Task{Name: name, Description: description[i], Id: id[i]})
		}

		data := DataUser{
			Tasks:    tasks,
			Username: cookieUsername,
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"withFields": data,
		})
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

		err = tmpl.Execute(w, map[string]interface{}{
			"withFields": data,
		})
		if err != nil {
			log.Fatalf("COULD NOT EXECUTE %v", err)
			http.Error(w, err.Error(), 400)
			return
		}
	}

}

func (h *handler) AddTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("Handling AddDescriptionTask request with parameters: %v", p)

	path := filepath.Join("public", "index2.html")
	tmpl, _ := template.ParseFiles(path)

	descriptionName := r.FormValue("descriptionName")
	description := r.FormValue("description")

	err := r.ParseForm()
	if err != nil {
		log.Fatalf("FAILED TO PARSE FORM %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookieUserID, err := r.Cookie("id")
	if err != nil {
		log.Fatalf("failed to get cookie %v", err)
		return
	}
	cookieID := cookieUserID.Value

	task := &entity.Task{
		AccountId:       cookieID,
		NameTask:        descriptionName,
		DescriptionTask: description,
	}

	id, err := h.service.CreateTask(r.Context(), task)
	if err != nil {
		fmt.Printf("failed to add taks %v", err)
		return
	}

	task.Id = id
	fmt.Println(task.Id)
	err = tmpl.Execute(w, map[string]interface{}{
		"taskID": task,
	})
	if err != nil {
		log.Fatalf("COULD NOT EXECUTE %v", err)
		http.Error(w, err.Error(), 400)
		return
	}
}

func (h *handler) DeleteTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("Handling DeleteDescriptionTask request with parameters: %v", p)

	id := r.FormValue("taskId")

	taskID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("failed to converti in int %v", nil)
		return
	}
	task := &entity.Task{
		Id: taskID,
	}

	h.service.DeleteTask(r.Context(), task)

}
