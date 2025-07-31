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
	default:
		fmt.Println("Unknown command:", command[1])
	}

}

func IsValidCommandInput(input string) bool {
	if input == "" || input == " " {
		return false
	}
	inputParts := strings.Split(input, " ")
	if inputParts[0] != "task-cli" {
		fmt.Println("The instruction must start with 'task-cli'")
		return false
	} else if len(inputParts) < 2 {
		fmt.Println("The instruction must have a command")
		return false
	} else if !IsValidCommand(strings.TrimSpace(inputParts[1])) {
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
	fmt.Println("Type the instruction preceded by 'task-cli'")
	fmt.Println("Type 'terminate' to exit")
	userInput := ""
	for {
		userInput = GetUserInput()

		if !IsValidCommandInput(userInput) {
			fmt.Println("Invalid input")
			continue
		}
		userCommand := strings.Split(userInput, " ")
		command := strings.TrimSpace(userCommand[1])
		ExecuteCommand(command, &userInput, &tasks)
		userInput = ""
		if command == "terminate" {
			fmt.Println("Goodbye")
			break
		}
	}

}
