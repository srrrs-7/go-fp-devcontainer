package tasks

import (
	"context"
	"database/sql"
	"utils/db/db"

	"github.com/google/uuid"
)

type MockQuerier struct {
	CountTasksByStatusFunc       func(ctx context.Context, status string) (int64, error)
	CountTasksByUserFunc         func(ctx context.Context, userID uuid.NullUUID) (int64, error)
	CreateTaskFunc               func(ctx context.Context, arg db.CreateTaskParams) (db.Task, error)
	DeleteTaskFunc               func(ctx context.Context, id uuid.UUID) error
	GetTaskFunc                  func(ctx context.Context, id uuid.UUID) (db.Task, error)
	ListOverdueTasksFunc         func(ctx context.Context) ([]db.Task, error)
	ListTasksFunc                func(ctx context.Context) ([]db.Task, error)
	ListTasksByStatusFunc        func(ctx context.Context, status string) ([]db.Task, error)
	ListTasksByUserFunc          func(ctx context.Context, userID uuid.NullUUID) ([]db.Task, error)
	ListTasksByUserAndStatusFunc func(ctx context.Context, arg db.ListTasksByUserAndStatusParams) ([]db.Task, error)
	ListUpcomingTasksFunc        func(ctx context.Context, dueDate sql.NullTime) ([]db.Task, error)
	UpdateTaskFunc               func(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error)
	UpdateTaskStatusFunc         func(ctx context.Context, arg db.UpdateTaskStatusParams) (db.Task, error)
}

func (m *MockQuerier) CountTasksByStatus(ctx context.Context, status string) (int64, error) {
	if m.CountTasksByStatusFunc != nil {
		return m.CountTasksByStatusFunc(ctx, status)
	}
	return 0, nil
}

func (m *MockQuerier) CountTasksByUser(ctx context.Context, userID uuid.NullUUID) (int64, error) {
	if m.CountTasksByUserFunc != nil {
		return m.CountTasksByUserFunc(ctx, userID)
	}
	return 0, nil
}

func (m *MockQuerier) CreateTask(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
	if m.CreateTaskFunc != nil {
		return m.CreateTaskFunc(ctx, arg)
	}
	return db.Task{}, nil
}

func (m *MockQuerier) DeleteTask(ctx context.Context, id uuid.UUID) error {
	if m.DeleteTaskFunc != nil {
		return m.DeleteTaskFunc(ctx, id)
	}
	return nil
}

func (m *MockQuerier) GetTask(ctx context.Context, id uuid.UUID) (db.Task, error) {
	if m.GetTaskFunc != nil {
		return m.GetTaskFunc(ctx, id)
	}
	return db.Task{}, nil
}

func (m *MockQuerier) ListOverdueTasks(ctx context.Context) ([]db.Task, error) {
	if m.ListOverdueTasksFunc != nil {
		return m.ListOverdueTasksFunc(ctx)
	}
	return nil, nil
}

func (m *MockQuerier) ListTasks(ctx context.Context) ([]db.Task, error) {
	if m.ListTasksFunc != nil {
		return m.ListTasksFunc(ctx)
	}
	return nil, nil
}

func (m *MockQuerier) ListTasksByStatus(ctx context.Context, status string) ([]db.Task, error) {
	if m.ListTasksByStatusFunc != nil {
		return m.ListTasksByStatusFunc(ctx, status)
	}
	return nil, nil
}

func (m *MockQuerier) ListTasksByUser(ctx context.Context, userID uuid.NullUUID) ([]db.Task, error) {
	if m.ListTasksByUserFunc != nil {
		return m.ListTasksByUserFunc(ctx, userID)
	}
	return nil, nil
}

func (m *MockQuerier) ListTasksByUserAndStatus(ctx context.Context, arg db.ListTasksByUserAndStatusParams) ([]db.Task, error) {
	if m.ListTasksByUserAndStatusFunc != nil {
		return m.ListTasksByUserAndStatusFunc(ctx, arg)
	}
	return nil, nil
}

func (m *MockQuerier) ListUpcomingTasks(ctx context.Context, dueDate sql.NullTime) ([]db.Task, error) {
	if m.ListUpcomingTasksFunc != nil {
		return m.ListUpcomingTasksFunc(ctx, dueDate)
	}
	return nil, nil
}

func (m *MockQuerier) UpdateTask(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error) {
	if m.UpdateTaskFunc != nil {
		return m.UpdateTaskFunc(ctx, arg)
	}
	return db.Task{}, nil
}

func (m *MockQuerier) UpdateTaskStatus(ctx context.Context, arg db.UpdateTaskStatusParams) (db.Task, error) {
	if m.UpdateTaskStatusFunc != nil {
		return m.UpdateTaskStatusFunc(ctx, arg)
	}
	return db.Task{}, nil
}
