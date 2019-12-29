package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
     r := RegisterHandlers()
     http.ListenAndServe(":8000", r)
//   listen->RegisterHandelrs->handlers
//   每个handler都是用不同的goroutine处理的，goroutine是非常轻量级的类线程的协程，
//   每个goroutine占用4kb，一瞬间可以创建几千几百个goroutine
}