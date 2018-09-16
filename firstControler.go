package main

import (
	"net/http"
	"fmt"
)

var firstController = FirstController{BaseController{"FirstOne"}}

func firstRoutes(h *Handler) {
	h.mux.HandleFunc("/hello/", firstController.sayHello)
	h.mux.HandleFunc("/fetch-task/", firstController.fetchTask)
}

// FirstController --
type FirstController struct {
	BaseController
}

func (c *FirstController) sayHello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Helllo!"))
}

func (c *FirstController) fetchTask(res http.ResponseWriter, req *http.Request) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}

	task1 := Task{}
	c.toStruct(response.Body, &task1)
	fmt.Println(task1)
	response.Body.Close()
	res.Write([]byte("Done"))
}

// Task --
type Task struct {
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}