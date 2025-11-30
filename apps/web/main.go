//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"syscall/js"
)

const apiBaseURL = "http://localhost:8080/api/v1"

func main() {
	// Register JavaScript functions
	js.Global().Set("fetchTasks", js.FuncOf(fetchTasks))
	js.Global().Set("addTask", js.FuncOf(addTask))

	// Initialize the UI
	renderApp()

	// Keep the program running
	select {}
}

// renderApp renders the initial HTML structure
func renderApp() {
	document := js.Global().Get("document")
	app := document.Call("getElementById", "app")

	html := `
		<div class="container">
			<h1>📝 Task Manager</h1>
			<p class="subtitle">Go + WebAssembly Demo</p>

			<div class="add-task">
				<input type="text" id="taskTitle" placeholder="Task title" />
				<input type="text" id="taskDesc" placeholder="Description (optional)" />
				<button onclick="addTask()">Add Task</button>
			</div>

			<div class="actions">
				<button onclick="fetchTasks()">🔄 Refresh Tasks</button>
			</div>

			<div id="tasks-container">
				<p class="hint">Click "Refresh Tasks" to load tasks from API</p>
			</div>

			<div id="status" class="status"></div>
		</div>
	`
	app.Set("innerHTML", html)
}

// Task represents a task from the API
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// TasksResponse represents the API response
type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}

// fetchTasks fetches tasks from the API
func fetchTasks(this js.Value, args []js.Value) any {
	go func() {
		setStatus("Loading tasks...")

		resp, err := http.Get(apiBaseURL + "/tasks")
		if err != nil {
			setStatus(fmt.Sprintf("❌ Error: %v", err))
			return
		}
		defer resp.Body.Close()

		var tasksResp TasksResponse
		if err := json.NewDecoder(resp.Body).Decode(&tasksResp); err != nil {
			setStatus(fmt.Sprintf("❌ Parse error: %v", err))
			return
		}

		renderTasks(tasksResp.Tasks)
		setStatus(fmt.Sprintf("✅ Loaded %d tasks", len(tasksResp.Tasks)))
	}()

	return nil
}

// addTask adds a new task via API
func addTask(this js.Value, args []js.Value) any {
	go func() {
		document := js.Global().Get("document")
		titleInput := document.Call("getElementById", "taskTitle")
		descInput := document.Call("getElementById", "taskDesc")

		title := titleInput.Get("value").String()
		desc := descInput.Get("value").String()

		if title == "" {
			setStatus("❌ Title is required")
			return
		}

		setStatus("Adding task...")

		// Create request body
		body := fmt.Sprintf(`{"title":"%s","description":"%s"}`, title, desc)

		resp, err := http.Post(
			apiBaseURL+"/tasks",
			"application/json",
			stringReader(body),
		)
		if err != nil {
			setStatus(fmt.Sprintf("❌ Error: %v", err))
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
			setStatus("✅ Task added!")
			// Clear inputs
			titleInput.Set("value", "")
			descInput.Set("value", "")
			// Refresh task list
			fetchTasks(js.Value{}, nil)
		} else {
			setStatus(fmt.Sprintf("❌ Failed: status %d", resp.StatusCode))
		}
	}()

	return nil
}

// renderTasks renders the task list
func renderTasks(tasks []Task) {
	document := js.Global().Get("document")
	container := document.Call("getElementById", "tasks-container")

	if len(tasks) == 0 {
		container.Set("innerHTML", `<p class="empty">No tasks found</p>`)
		return
	}

	html := `<ul class="task-list">`
	for _, task := range tasks {
		completedClass := ""
		completedIcon := "⬜"
		if task.Completed {
			completedClass = "completed"
			completedIcon = "✅"
		}

		html += fmt.Sprintf(`
			<li class="task-item %s">
				<span class="task-status">%s</span>
				<div class="task-content">
					<strong>%s</strong>
					<span class="task-desc">%s</span>
				</div>
			</li>
		`, completedClass, completedIcon, task.Title, task.Description)
	}
	html += `</ul>`

	container.Set("innerHTML", html)
}

// setStatus updates the status message
func setStatus(msg string) {
	document := js.Global().Get("document")
	status := document.Call("getElementById", "status")
	status.Set("innerHTML", msg)
}

// stringReader implements io.Reader for a string
type stringReaderImpl struct {
	s string
	i int
}

func stringReader(s string) *stringReaderImpl {
	return &stringReaderImpl{s: s}
}

func (r *stringReaderImpl) Read(p []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, fmt.Errorf("EOF")
	}
	n = copy(p, r.s[r.i:])
	r.i += n
	return n, nil
}
