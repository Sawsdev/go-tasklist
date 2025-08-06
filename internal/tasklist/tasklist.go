package tasklist

import (
	"fmt"
	"time"
	"github.com/sawsdevx8/tasktracker/internal/task"
)

const (
	layout = "2006-01-02 15:04:05"
)

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

func printTask(task task.Task)  {
	fmt.Println("-------****--------")
		fmt.Println("Id:", task.Id)
		fmt.Println("Description:", task.Description)
		fmt.Println("Status:", task.Status)
		fmt.Println("Created at:", task.CreatedAt)
		fmt.Println("Updated at:", task.UpdatedAt)
		fmt.Println("-------****--------")
}
