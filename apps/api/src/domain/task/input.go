package task

// TaskCmd represents a command to create or update a task.
// It contains only the mutable properties of a task (excluding ID).
type TaskCmd struct {
	Title       TaskTitle
	Description TaskDescription
	Status      TaskStatus
}
