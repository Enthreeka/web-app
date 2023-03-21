package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
	"web/internal/account/db"
	http3 "web/internal/account/delivery/http"
	usecase2 "web/internal/account/usecase"
	"web/internal/config"
	"web/internal/entity"
	db2 "web/internal/user/db"
	http2 "web/internal/user/delivery/http"
	"web/internal/user/usecase"
	"web/pkg/postgresql"
)

func main() {
	cfg := config.GetConfig()

	dataBase, err := postgresql.NewConnect(context.TODO(), cfg.Storage)
	if err != nil {
		log.Fatalf("%W failed to init DB connection", err)
	}

	userRepository := db2.NewUserRepository(dataBase)
	accountRepository := db.NewAccountRepository(dataBase)

	userService := usecase.NewUserService(userRepository)
	accountService := usecase2.NewAccountService(accountRepository)

	router := httprouter.New()

	user := &entity.User{}
	userHandler := http2.NewHandler(userService, user)
	accountHandler := http3.NewAccountHandler(accountService)

	userHandler.Register(router)
	accountHandler.Register(router)

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
