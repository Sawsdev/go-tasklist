package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sawsdevx8/tasktracker/internal/tasklist"
)

const (
	ADD          = "add"
	DELETE       = "delete"
	LIST         = "list"
	UPDATE       = "update"
	MARKDONE     = "mark-done"
	MARKPROGRESS = "mark-progress"
	TERMINATE    = "terminate"
)

var cliCommands = map[string]string{
	ADD:          "add",
	DELETE:       "delete",
	LIST:         "list",
	UPDATE:       "update",
	MARKDONE:     "mark-done",
	MARKPROGRESS: "mark-progress",
	TERMINATE:    "terminate",
}

func IsValidCommand(command string) bool {
	for _, c := range cliCommands {
		if c == command {
			return true
		}
	}
	return false
}

func GetCommandDescription(command string) string {
	switch command {
	case ADD:
		return "Add a task to the list \nej: add \"Make coffee\""
	case DELETE:
		return "Delete a task from the list \nej: delete 1"
	case LIST:
		return "List all tasks"
	case UPDATE:
		return "Update a task \nej: update 1 'Make coffee with milk'"
	case MARKDONE:
		return "Mark a task as done \nej: mark-done 1"
	case MARKPROGRESS:
		return "Mark a task as in progress \nej: mark-progress 1"
	case TERMINATE:
		return "Terminate the program: terminate"
	default:
		return fmt.Sprintf("Unknown command: %s", command)
	}
}

func GetDescriptionFromInput(input string) string {
	description := strings.Split(input, "\"")
	return description[1]
}

func ExecuteCommand(command string, userInput *string, tasks *tasklist.TaskList) {
	fmt.Println("Executing command:", command)
	cleanCommand := strings.TrimSpace(command)
	switch cleanCommand {
	case ADD:
		if !strings.Contains(*userInput, "\"") {
			fmt.Println("Invalid input, no correct description found")
			fmt.Println("ej: add \"Make coffee\"")
			return
		}
		tasklist.AddNewTask(tasks, GetDescriptionFromInput(*userInput))
	case DELETE:
		// tasklist.DeleteTask(&tasklist.TaskList, command[2])
	case LIST:
		tasklist.ShowTaskList(tasks)
	case UPDATE:
		// tasklist.UpdateTask(&tasklist.TaskList, command[2], command[3])
	case MARKDONE:
		// tasklist.MarkTaskAsDone(&tasklist.TaskList, command[2])
	case MARKPROGRESS:
	// tasklist.MarkTaskAsInProgress(&tasklist.TaskList, command[2])
	case TERMINATE:
		fmt.Println("Exiting program")
	default:
		fmt.Println("Unknown command:", cleanCommand)
	}

}

func IsValidCommandInput(input string) bool {
	if input == "" || input == " " {
		return false
	}
	inputParts := strings.Split(input, " ")
	if len(inputParts) < 1 {
		fmt.Println("The instruction must have a command")
		return false
	} else if !IsValidCommand(strings.TrimSpace(inputParts[0])) {
		fmt.Println("Invalid command")
		return false
	}

	return true
}


func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func Start() {

	tasks := tasklist.NewTaskList()
	fmt.Println("Task CLI")
	fmt.Println("Available commands:")
	for command, description := range cliCommands {
		fmt.Println(command, ":")
		fmt.Println(GetCommandDescription(description))
		fmt.Println("----------------")
	}
	fmt.Println("Type 'terminate' to exit")
	userInput := ""
	for {
		fmt.Print("task-cli ")
		userInput = GetUserInput()
		fmt.Println(userInput)

		if !IsValidCommandInput(userInput) {
			continue
		}
		userCommand := strings.Split(userInput, " ")
		command := strings.TrimSpace(userCommand[0])
		ExecuteCommand(command, &userInput, &tasks)
		userInput = ""
		if command == "terminate" {
			fmt.Println("Goodbye")
			break
		}
	}

}
