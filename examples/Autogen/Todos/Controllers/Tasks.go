package pgucontrollers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	SQLL "github.com/eshu0/persist/pkg/sqllite"	
	per "github.com/eshu0/persist/pkg/interfaces"
	data "github.com/eshu0/Pangu/examples/Autogen/Todos/Models"
)

// Controller

type TasksController struct {
	TasksHandler *TasksHandler // Storage handler
}

func NewTasksController(handler *TasksHandler) *TasksController {
	ds := TasksController{}
	ds.TasksHandler = handler
	return &ds
}

func (controller *TasksController) HandleCreateRequest(DataIn Task) Task {
	result := controller.TasksHandler.Create(DataIn)
	fmt.Println("----")
	fmt.Println("Result")
	fmt.Println("----")
	fmt.Println(result)
	return result
}


