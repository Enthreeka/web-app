package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
	"web/internal/config"
	"web/internal/user"
	db2 "web/internal/user/db"
	"web/pkg/postgresql"
)

func main() {

	cfg := config.GetConfig()

	db, err := postgresql.NewConnect(context.TODO(), cfg.Storage)
	if err != nil {
		log.Fatalf("%W failed to init DB connection", err)
	}

	userRepository := db2.NewUserRepository(db)

	service := user.NewService(userRepository)

	router := httprouter.New()

	handler := user.NewHandler(service)
	handler.Register(router)

	//user := &entity.User{
	//	Login:    "olegLeon",
	//	Password: "daun",
	//}
	//userRepository.CreateUser(context.TODO(), user)

	//accountRepository := db3.NewAccountRepository(db)
	//
	//account := &entity.Account{
	//	UserId: "4",
	//	Name:   "ilyha",
	//}
	//
	//err = accountRepository.CreateAccount(context.TODO(), account)
	//if err != nil {
	//	log.Fatalf("failed to create account %v", err)
	//}

	//err = accountRepository.UpdateName(context.TODO(), account, 1)
	//if err != nil {
	//	log.Fatalf("failed to update account error - %v", err)
	//}

	//user, err := userRepository.GetUser(context.TODO(), 6)
	//if err != nil {
	//	log.Fatalf("failed to get user -  %v", err)
	//}
	//fmt.Println(user)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Server.Hostname, cfg.Server.Port))
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(server.Serve(listener))

}
