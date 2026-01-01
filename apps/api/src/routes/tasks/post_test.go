package tasks

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"utils/db/db"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestPostHandler(t *testing.T) {
	t.Run("201 Created", func(t *testing.T) {
		type expected struct {
			statusCode int
			hasID      bool
		}

		tests := []struct {
			name     string
			formData map[string]string
			expected expected
		}{
			{
				name: "valid request with title and description",
				formData: map[string]string{
					"title":       "New Task",
					"description": "Task Description",
				},
				expected: expected{
					statusCode: http.StatusCreated,
					hasID:      true,
				},
			},
			{
				name: "valid request with title only",
				formData: map[string]string{
					"title": "New Task",
				},
				expected: expected{
					statusCode: http.StatusCreated,
					hasID:      true,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				formData := url.Values{}
				for k, v := range tt.formData {
					formData.Set(k, v)
				}
				req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					CreateTaskFunc: func(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
						return db.Task{
							ID:          uuid.New(),
							Title:       arg.Title,
							Description: arg.Description,
							Status:      "pending",
						}, nil
					},
				}

				handler := NewPostHandler(mockDB)
				handler.ServeHTTP(w, req)

				resp := w.Result()
				if diff := cmp.Diff(tt.expected.statusCode, resp.StatusCode); diff != "" {
					t.Errorf("status code mismatch (-want +got):\n%s", diff)
				}

				var result map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				_, hasID := result["id"]
				if diff := cmp.Diff(tt.expected.hasID, hasID); diff != "" {
					t.Errorf("'id' field presence mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})

	t.Run("400 Bad Request", func(t *testing.T) {
		type expected struct {
			statusCode int
			hasType    bool
		}

		tests := []struct {
			name     string
			formData map[string]string
			expected expected
		}{
			{
				name: "title too short",
				formData: map[string]string{
					"title": "ab",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
			{
				name: "missing title",
				formData: map[string]string{
					"description": "Only description",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
			{
				name: "empty title",
				formData: map[string]string{
					"title": "",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				formData := url.Values{}
				for k, v := range tt.formData {
					formData.Set(k, v)
				}
				req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					CreateTaskFunc: func(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
						return db.Task{
							ID:          uuid.New(),
							Title:       arg.Title,
							Description: arg.Description,
							Status:      "pending",
						}, nil
					},
				}

				handler := NewPostHandler(mockDB)
				handler.ServeHTTP(w, req)

				resp := w.Result()
				if diff := cmp.Diff(tt.expected.statusCode, resp.StatusCode); diff != "" {
					t.Errorf("status code mismatch (-want +got):\n%s", diff)
				}

				var result map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				_, hasType := result["type"]
				if diff := cmp.Diff(tt.expected.hasType, hasType); diff != "" {
					t.Errorf("'type' field presence mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})
}
