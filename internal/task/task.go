package task


type Task struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}



func NewTask(id int, description string , createdAt string, updatedAt string) Task {
	return Task{
		Id: id,
		Description: description,
		Status: "todo",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}