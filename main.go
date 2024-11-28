package main

import (
	account "api-login-proto/account"
	"service-account/controllers"
	"standard-library/grpc"
	initilize "standard-library/initialize"
	"standard-library/validation"
)

func main() {
	initilize.InitLogs()
	initilize.InitNacosConfig()
	initilize.InitRedis()
	initilize.InitDB()
	initilize.InitMail()
	validation.Init()

	srv := grpc.NewServer()
	account.RegisterUserAccountServiceServer(srv.Srv, &controllers.AccountController{})
	initilize.RunGRPC(srv)
}
