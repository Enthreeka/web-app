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
	router.POST(SaveImage, h.ImageSaveHandler)
}

func getUserID(r *http.Request) string {
	cookieUserID, err := r.Cookie("id")
	if err != nil {
		log.Fatalf("failed to get cookie %v", err)
		return ""
	}
	return cookieUserID.Value
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
	cookieUsername := cookieUN.Value

	userID := getUserID(r)

	nameAccount, err := h.service.GetName(r.Context(), userID)
	if err != nil {
		log.Fatalf("failed to get name %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type InfoAccount struct {
		Name string
	}

	id, name, description, err := h.service.GetTasks(r.Context(), userID)
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

		imgSrc, err := h.service.GetPhoto(r.Context(), userID)
		if err != nil {
			log.Printf("failed to get photo ERROR: %v", err)
		}
		img := template.URL(imgSrc)

		dataImg := struct {
			ImageSrc interface{}
		}{
			img,
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"withFields":         data,
			"inforamtionAccount": dataName,
			"image":              dataImg,
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
		imgSrc, err := h.service.GetPhoto(r.Context(), userID)
		if err != nil {
			log.Printf("failed to get photo ERROR: %v", err)
		}
		img := template.URL(imgSrc)
		dataImg := struct {
			ImageSrc interface{}
		}{
			img,
		}

		err = tmpl.Execute(w, map[string]interface{}{
			"withFields":         data,
			"inforamtionAccount": dataName,
			"image":              dataImg,
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

	userID := getUserID(r)

	task := &entity.Task{
		AccountId:       userID,
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

	userID := getUserID(r)
	err := h.service.Leave(r.Context(), userID)
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

		userID := getUserID(r)
		err = h.service.SaveName(r.Context(), userID, data.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *handler) ImageSaveHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("Handling ImageSaveHandler request")

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

	userID := getUserID(r)
	err = h.service.AddPhoto(r.Context(), userID, imgBytes)
	if err != nil {
		log.Println("failed to add photo %v", err)
		return
	}

}

// ImageLoadingHandler
func (h *handler) ImageLoadingHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	path := filepath.Join("public", "index2.html")
	tmpl, _ := template.ParseFiles(path)

	userID := getUserID(r)
	imgSrc, err := h.service.GetPhoto(r.Context(), userID)
	if err != nil {
		log.Printf("failed to get photo ERROR: %v", err)
	}

	data := struct {
		ImageSrc string
	}{
		imgSrc,
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"image": data,
	})
	if err != nil {
		log.Fatalf("COULD NOT EXECUTE %v", err)
		http.Error(w, err.Error(), 400)
		return
	}
}
