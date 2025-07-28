package task


type Task struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

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

func NewTask(id int, description string , createdAt string, updatedAt string) Task {
	return Task{
		Id: id,
		Description: description,
		Status: TODO,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}