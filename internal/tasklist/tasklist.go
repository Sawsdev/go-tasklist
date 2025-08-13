package tasklist

import (
	"fmt"
	"time"
	"github.com/sawsdevx8/tasktracker/internal/task"
	"slices"
)

const (
	layout = "2006-01-02 15:04:05"
)

const (
	TODO = "todo"
	DONE = "done"
	PROGRESS = "in-progress"
)


var Statuses = map[string]string{
	TODO: "todo",
	DONE : "done", 
	PROGRESS : "in-progress",
}

type TaskList struct {
	Tasks []task.Task
}

func NewTaskList() TaskList {
	return TaskList{
		Tasks: []task.Task{},
	}
}

func AddNewTask(taskList *TaskList, description string) {
	task := task.NewTask(
		len(taskList.Tasks) + 1,
		description,
		time.Now().Format(layout),
		time.Now().Format(layout))
	taskList.Tasks = append(taskList.Tasks, task)
	fmt.Println("Task added \n", task)
}

func ShowTaskList(taskList *TaskList, status string) {
	fmt.Println("All task in the list:")
	for _, task := range taskList.Tasks {
		if status == "all"{
			printTask(task)
		} else if task.Status == status {
			printTask(task)
		}
	}
}

func UpdateTask(taskList *TaskList, id int, description string, status string) {
	updateTask(taskList, id, description, status)
}

func MarkTaskAsDone(taskList *TaskList, id int) {
	updateTask(taskList, id, "", DONE)
}

func MarkTaskAsInProgress(taskList *TaskList, id int) {
	updateTask(taskList, id, "", PROGRESS)
}

func DeleteTask(taskList *TaskList, id int) {
	for i := range taskList.Tasks {
		if taskList.Tasks[i].Id == id {
			taskList.Tasks = slices.DeleteFunc(taskList.Tasks, func(t task.Task) bool {
				return t.Id == id
			})
			fmt.Println("Task deleted")
			return
		}
	}
	fmt.Println("Task not found")
}


func updateTask(taskList *TaskList, id int, description string, status string) {
	for i := range taskList.Tasks {
		if taskList.Tasks[i].Id == id {
			
			if status == "" && description == "" {
				fmt.Println("Nothing to update")
				return
			}
			if description != "" {
				taskList.Tasks[i].Description = description
			} else if condition := Statuses[status]; condition != "" {
				taskList.Tasks[i].Status = status
			}
			taskList.Tasks[i].UpdatedAt = time.Now().Format(layout)
			fmt.Println("Task updated \n", taskList.Tasks[i])
			return
		}
	}
	fmt.Println("Task not found")
}


func printTask(task task.Task)  {
	fmt.Println("-------****--------")
		fmt.Println("Id:", task.Id)
		fmt.Println("Description:", task.Description)
		fmt.Println("Status:", task.Status)
		fmt.Println("Created at:", task.CreatedAt)
		fmt.Println("Updated at:", task.UpdatedAt)
		fmt.Println("-------****--------")
}
