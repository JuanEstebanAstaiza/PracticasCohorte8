package main

//ejercicio 2

import (
	"errors"
	"fmt"
)

type TaskStatus string

const (
	Pending    TaskStatus = "pendiente"
	InProgress TaskStatus = "en progreso"
	Completed  TaskStatus = "completada"
)

type Task struct {
	Name        string
	Description string
	Assignee    string
	Status      TaskStatus
}

type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) AddTask(name, description string) {
	task := Task{

		Name:        name,
		Description: description,
		Status:      Pending,
	}
	tm.tasks = append(tm.tasks, task)
}

func (tm *TaskManager) AssignTask(name, assignee string) error {
	for i, task := range tm.tasks {
		if task.Name == name {
			tm.tasks[i].Assignee = assignee
			return nil
		}
	}
	return errors.New("tarea no encontrada")
}

func (tm *TaskManager) UpdateTaskStatus(name string, status TaskStatus) error {
	for i, task := range tm.tasks {
		if task.Name == name {
			tm.tasks[i].Status = status
			return nil
		}
	}
	return errors.New("Tarea no encontrada")
}

func (tm *TaskManager) PendingTasks() []Task {
	var pendingTasks []Task
	for _, task := range tm.tasks {
		if task.Status == Pending {
			pendingTasks = append(pendingTasks, task)
		}
	}
	return pendingTasks
}

func main() {
	tm := &TaskManager{}

	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	tm.AddTask("Tarea 2", "Descripcion de la tarea 2")

	err := tm.AssignTask("Tarea 1", "Juan")
	if err != nil {
		fmt.Println("Error actualizando estado de tarea: ", err)
	}

	pendingTasks := tm.PendingTasks()
	fmt.Printf("Tareas pendientes: ")
	for _, task := range pendingTasks {
		fmt.Printf("Nombre: %s, Descripcion: %s, Responsable: %s, Estado: %s\n", task.Name, task.Description, task.Assignee, task.Status)
	}
}
