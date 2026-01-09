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
	"utils/testutil"

	"github.com/google/go-cmp/cmp"
)

func TestPutHandler(t *testing.T) {

	t.Run("200 OK", func(t *testing.T) {
		type expected struct {
			statusCode  int
			title       string
			description string
			status      string
		}

		tests := []struct {
			name     string
			setup    func(t *testing.T, q db.Querier) string
			formData func(taskID string) map[string]string
			expected expected
		}{
			{
				name: "update task title and status",
				setup: func(t *testing.T, q db.Querier) string {
					task, err := q.CreateTask(context.Background(), db.CreateTaskParams{
						Title:    "Original Title",
						Status:   "pending",
						Priority: "medium",
					})
					if err != nil {
						t.Fatalf("failed to create test task: %v", err)
					}
					return task.ID.String()
				},
				formData: func(taskID string) map[string]string {
					return map[string]string{
						"id":          taskID,
						"title":       "Updated Title",
						"description": "Updated Description",
						"status":      "completed",
					}
				},
				expected: expected{
					statusCode:  http.StatusOK,
					title:       "Updated Title",
					description: "Updated Description",
					status:      "completed",
				},
			},
			{
				name: "update task without description",
				setup: func(t *testing.T, q db.Querier) string {
					task, err := q.CreateTask(context.Background(), db.CreateTaskParams{
						Title:    "Original Title",
						Status:   "pending",
						Priority: "medium",
					})
					if err != nil {
						t.Fatalf("failed to create test task: %v", err)
					}
					return task.ID.String()
				},
				formData: func(taskID string) map[string]string {
					return map[string]string{
						"id":     taskID,
						"title":  "Updated Title Only",
						"status": "in_progress",
					}
				},
				expected: expected{
					statusCode:  http.StatusOK,
					title:       "Updated Title Only",
					description: "",
					status:      "in_progress",
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := testutil.SetupTestTx(t)
				taskID := tt.setup(t, q)

				formData := url.Values{}
				for k, v := range tt.formData(taskID) {
					formData.Set(k, v)
				}
				req := httptest.NewRequest(http.MethodPut, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				testutil.SetAuthHeader(req)

				w := httptest.NewRecorder()

				handler := NewPutHandler(q)
				handler.ServeHTTP(w, req)

				resp := w.Result()
				if diff := cmp.Diff(tt.expected.statusCode, resp.StatusCode); diff != "" {
					t.Errorf("status code mismatch (-want +got):\n%s", diff)
				}

				var result map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if diff := cmp.Diff(taskID, result["id"]); diff != "" {
					t.Errorf("id mismatch (-want +got):\n%s", diff)
				}

				if diff := cmp.Diff(tt.expected.title, result["title"]); diff != "" {
					t.Errorf("title mismatch (-want +got):\n%s", diff)
				}

				if diff := cmp.Diff(tt.expected.description, result["description"]); diff != "" {
					t.Errorf("description mismatch (-want +got):\n%s", diff)
				}

				if diff := cmp.Diff(tt.expected.status, result["status"]); diff != "" {
					t.Errorf("status mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})

	t.Run("404 Not Found", func(t *testing.T) {
		tests := []struct {
			name           string
			formData       map[string]string
			expectedStatus int
		}{
			{
				name: "update non-existent task",
				formData: map[string]string{
					"id":     "00000000-0000-0000-0000-000000000000",
					"title":  "Updated Title",
					"status": "pending",
				},
				expectedStatus: http.StatusNotFound,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				q := testutil.SetupTestTx(t)

				formData := url.Values{}
				for k, v := range tt.formData {
					formData.Set(k, v)
				}
				req := httptest.NewRequest(http.MethodPut, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				testutil.SetAuthHeader(req)

				w := httptest.NewRecorder()

				handler := NewPutHandler(q)
				handler.ServeHTTP(w, req)

				resp := w.Result()
				if diff := cmp.Diff(tt.expectedStatus, resp.StatusCode); diff != "" {
					t.Errorf("status code mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})
}
