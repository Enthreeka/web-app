package http

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"web/internal/account/usecase"
	"web/internal/apperror"
	"web/internal/entity"
	"web/internal/handlers"
	"web/internal/validation"
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

const (
	startPage = "/"
	dashboard = "/dashboard"
	add       = "/dashboard/add"
	delete    = "/dashboard/delete"
	logout    = "/dashboard/leave"
	edit      = "/dashboard/edit"
	saveName  = "/dashboard/update/name"
	SaveImage = "/dashboard/image"
)

func (h *handler) Register(router *httprouter.Router) {
	log.Println("Registering account routes...")

	router.GET(dashboard, apperror.AuthMiddleware(h.GetTask))
	router.POST(add, h.AddTask)
	router.DELETE(delete, h.DeleteTask)
	router.POST(logout, h.logoutHandler)
	router.PUT(edit, h.EditHandler)
	router.POST(saveName, h.SaveNameHandler)
	router.POST(SaveImage, h.ImageHandler)
}

func (h *handler) GetTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("Handling GetTasks request with parameters: %v", p)
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

	nameAccount, err := h.service.GetName(r.Context(), cookieID)
	if err != nil {
		log.Fatalf("failed to get name %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type InfoAccount struct {
		Name string
	}

	id, name, description, err := h.service.GetTasks(r.Context(), cookieID)
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

		dataName := InfoAccount{
			Name: nameAccount,
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"withFields":         data,
			"inforamtionAccount": dataName,
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
			Name     string
		}
		data := DataUser{
			Username: cookieUsername,
		}
		dataName := InfoAccount{
			Name: nameAccount,
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"withFields":         data,
			"inforamtionAccount": dataName,
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

func (h *handler) logoutHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	log.Println("Handling logoutHandler request")

	cookieUserID, err := r.Cookie("id")
	if err != nil {
		log.Fatalf("failed to get cookie %v", err)
		return
	}
	cookieID := cookieUserID.Value

	err = h.service.Leave(r.Context(), cookieID)
	if err != nil {
		log.Fatalf("failed to set null to jwt token %v", err)
		return
	}

	cookie := &http.Cookie{
		Name:   "jwt",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, startPage, http.StatusSeeOther)
}

func (h *handler) EditHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Handling EditHandler request")

	//if r.Method == "PUT" {
	//	taskNameForm := r.FormValue("taskName")
	//	taskDescriptionForm := r.FormValue("taskDescription")
	//	//id := r.FormValue("taskId")
	//
	//	id := r.PostForm.Get("taskId")
	//	err := r.ParseForm()
	//	if err != nil {
	//		log.Fatalf("FAILED TO PARSE FORM %v", err)
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//		return
	//	}
	//	idInt, _ := strconv.Atoi(id)
	//
	//	taskNameDB, taskDescriptionDB, err := h.service.GetTask(r.Context(), idInt)
	//	if err != nil {
	//		log.Printf("failed to get task in handler %v", err)
	//		return
	//	}
	//
	//	if taskDescriptionForm != taskDescriptionDB {
	//		err = h.service.UpdateDescriptionTask(r.Context(), taskDescriptionForm, idInt)
	//		if err != nil {
	//			log.Printf("Failed to update description task %v", err)
	//			return
	//		}
	//	}
	//	if taskNameForm != taskNameDB {
	//		err = h.service.UpdateNameTask(r.Context(), taskDescriptionForm, idInt)
	//		if err != nil {
	//			log.Printf("Failed to update name task %v", err)
	//			return
	//		}
	//	}
	//}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}

func (h *handler) SaveNameHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Handling SaveNameHandler request")

	type Data struct {
		Value string `json:"value"`
	}
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	if !validation.IsValidationName(data.Value) {
		fmt.Println("name did not meet the requirements")
		return
	} else {

		cookieUserID, err := r.Cookie("id")
		if err != nil {
			log.Fatalf("failed to get cookie %v", err)
			return
		}
		userID := cookieUserID.Value

		err = h.service.SaveName(r.Context(), userID, data.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *handler) ImageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Handling ImageHandler request")

	err := r.ParseMultipartForm(32 << 20) // максимальный размер 32MB
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("failed ti get file from form %v", err)
		return
	}
	defer file.Close()

	imgBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	cookieUserID, err := r.Cookie("id")
	if err != nil {
		log.Fatalf("failed to get cookie %v", err)
		return
	}
	userID := cookieUserID.Value

	err = h.service.AddPhoto(r.Context(), userID, imgBytes)
	if err != nil {
		log.Println("failed to add photo %v", err)
		return
	}

	fmt.Println(imgBytes)

}
