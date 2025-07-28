package cli

import (
	"fmt"
	"strings"
	"os"
	"bufio"
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
		return "Add a task to the list ej: add 'Make coffee'"
	case DELETE:
		return "Delete a task from the list ej: delete 1"
	case LIST:
		return "List all tasks"
	case UPDATE:
		return "Update a task ej: update 1 'Make coffee with milk'"
	case MARKDONE:
		return "Mark a task as done ej: mark-done 1"
	case MARKPROGRESS:
		return "Mark a task as in progress"
	case TERMINATE:
		return "Terminate the program: terminate"
	default:
		return fmt.Sprintf("Unknown command: %s", command)
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

func GetCommand(input string) string {
	inputParts := strings.Split(input, " ")
	return inputParts[1]
}

func GetUserInput(completeText *string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	*completeText += text
}

func Start() {
	fmt.Println("Task CLI")
	fmt.Println("Available commands:")
	for command, description := range cliCommands {
		fmt.Printf("%s: %s\n", command, description)
	}
	fmt.Println("Type the instruction preceded by 'task-cli'")
	fmt.Println("Type 'terminate' to exit")
	userInput := ""
	for {
		GetUserInput(&userInput)
		
		if !IsValidCommandInput(userInput) {
			fmt.Println("Invalid input")
			continue
		}
		command := strings.TrimSpace(GetCommand(userInput))
		fmt.Println(GetCommandDescription(command))
		userInput = ""
		if command == "terminate" {
			fmt.Println("Goodbye")
			break
		}
	}
	
}


