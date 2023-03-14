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

func (h *handler) Register(router *httprouter.Router) {

	log.Println("Registering account routes...")

	//router.GET("/dashboard", apperror.AuthMiddleware(h.AccountPage))

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

	q := r.URL.Query()
	userID, _ := strconv.Atoi(q.Get("id"))

	name, description, err := h.service.GetTask(r.Context(), userID)
	if name != nil || description != nil {
		if err != nil {
			log.Printf("failed with get task in handler %v", err)
			return
		}

		type Task struct {
			Name        string
			Description string
		}

		tasks := []Task{}
		for i, name := range name {
			tasks = append(tasks, Task{Name: name, Description: description[i]})
		}

		fmt.Printf("tasks - %v \n", tasks)

		//user := entity.User{
		//	Login: "Login",
		//}
		//
		//err = tmpl.Execute(w, user)
		//if err != nil {
		//	log.Fatalf("COULD NOT EXECUTE %v", err)
		//	http.Error(w, err.Error(), 400)
		//	return
		//}

		err = tmpl.Execute(w, tasks)
		if err != nil {
			log.Fatalf("COULD NOT EXECUTE %v", err)
			http.Error(w, err.Error(), 400)
			return
		}
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("COULD NOT EXECUTE %v", err)
		http.Error(w, err.Error(), 400)
		return
	}
}

//func (h *handler) GetTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	log.Printf("Handling GetTask request with parameters: %v", p)
//
//	account := &entity.Account{UserId: 14}
//
//	tasks, err := h.service.GetTask(r.Context(), account)
//	if err != nil {
//		log.Printf("failed to get task in handler %v", err)
//		return
//	}
//
//	path := filepath.Join("public", "index2.html")
//	tmpl, err := template.ParseFiles(path)
//	if err != nil {
//		http.Error(w, err.Error(), 400)
//		return
//	}
//	//type TaskData struct {
//	//	Task []entity.Account
//	//}
//	//
//	//tasksData := TaskData{Task: tasks}
//	//
//	//tasksJSON, err := json.Marshal(tasksData)
//	//if err != nil {
//	//	log.Printf("failed to marshal tasks data: %v", err)
//	//	return
//	//}
//	//
//	//var data TaskData
//	//err = json.Unmarshal(tasksJSON, &data)
//	//if err != nil {
//	//	log.Printf("failed to unmarshal tasks data: %v", err)
//	//	return
//	//}
//	//
//	//err = tmpl.Execute(w, map[string]interface{}{
//	//	"users": data,
//	//})
//
//	err = tmpl.Execute(w, map[string]interface{}{
//		"tasks": tasks,
//	})
//	if err != nil {
//		log.Fatalf("COULD NOT EXECUTE %v", err)
//		http.Error(w, err.Error(), 400)
//		return
//	}
//
//}

func (h *handler) AddTask(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Header.Set("Content-Type", "application/json")
	log.Printf("Handling CreateTask request with parameters: %v", p)

	descriptionName := r.FormValue("descriptionName")
	description := r.FormValue("description")
	//userid := p.ByName("user_id")

	//fmt.Println(userid)
	fmt.Println("description = ", description)
	fmt.Println("descriptionName = ", descriptionName)

	//q := r.URL.Query()
	//q.Get("user_id")
	//query := q.Encode()
	//	totalString := strings.Trim(query, "user=")
	//fmt.Println(query)

	//userIdInt, _ := strconv.Atoi(userid)
	userIdInt := 14

	account := &entity.Account{
		UserId:          userIdInt,
		NameTask:        descriptionName,
		DescriptionTask: description,
	}

	err := h.service.AddTask(r.Context(), account)
	if err != nil {
		fmt.Printf("failed to add taks %v", err)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}
