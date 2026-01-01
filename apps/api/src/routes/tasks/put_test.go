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
)

func TestPutHandler(t *testing.T) {
	t.Run("200 OK", func(t *testing.T) {
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
				name: "valid request with all fields",
				formData: map[string]string{
					"id":          "550e8400-e29b-41d4-a716-446655440000",
					"title":       "Updated Task",
					"description": "Updated Description",
					"status":      "completed",
				},
				expected: expected{
					statusCode: http.StatusOK,
					hasID:      true,
				},
			},
			{
				name: "valid request without description",
				formData: map[string]string{
					"id":     "550e8400-e29b-41d4-a716-446655440000",
					"title":  "Updated Task",
					"status": "pending",
				},
				expected: expected{
					statusCode: http.StatusOK,
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
				req := httptest.NewRequest(http.MethodPut, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					UpdateTaskFunc: func(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error) {
						return db.Task{
							ID:          arg.ID,
							Title:       arg.Title.String,
							Description: arg.Description,
							Status:      arg.Status.String,
						}, nil
					},
				}

				handler := NewPutHandler(mockDB)
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
				name: "invalid uuid",
				formData: map[string]string{
					"id":    "invalid-uuid",
					"title": "Updated Task",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
			{
				name: "missing id",
				formData: map[string]string{
					"title": "Updated Task",
				},
				expected: expected{
					statusCode: http.StatusBadRequest,
					hasType:    true,
				},
			},
			{
				name: "title too short",
				formData: map[string]string{
					"id":    "00000000-0000-0000-0000-000000000001",
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
					"id": "00000000-0000-0000-0000-000000000001",
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
				req := httptest.NewRequest(http.MethodPut, "/tasks", strings.NewReader(formData.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				w := httptest.NewRecorder()
				mockDB := &MockQuerier{
					UpdateTaskFunc: func(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error) {
						return db.Task{
							ID:          arg.ID,
							Title:       arg.Title.String,
							Description: arg.Description,
							Status:      arg.Status.String,
						}, nil
					},
				}

				handler := NewPutHandler(mockDB)
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
