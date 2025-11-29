package tasks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListHandler(t *testing.T) {
	type args struct {
		queryParams map[string]string
	}
	type expected struct {
		statusCode int
		hasError   bool
	}

	tests := []struct {
		testName string
		args     args
		expected expected
	}{
		{
			testName: "valid request",
			args: args{
				queryParams: map[string]string{
					"id":          "550e8400-e29b-41d4-a716-446655440000",
					"title":       "Test Task",
					"description": "Test Description",
					"completed":   "false",
				},
			},
			expected: expected{
				statusCode: http.StatusOK,
				hasError:   false,
			},
		},
		{
			testName: "invalid uuid",
			args: args{
				queryParams: map[string]string{
					"id":    "invalid",
					"title": "Test Task",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				hasError:   true,
			},
		},
		{
			testName: "title too short",
			args: args{
				queryParams: map[string]string{
					"id":    "00000000-0000-0000-0000-000000000001",
					"title": "ab",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				hasError:   true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
			q := req.URL.Query()
			for k, v := range tt.args.queryParams {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()

			w := httptest.NewRecorder()
			ListHandler(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expected.statusCode {
				t.Errorf("expected status %v, got %v", tt.expected.statusCode, resp.StatusCode)
			}

			if resp.Header.Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type application/json, got %v", resp.Header.Get("Content-Type"))
			}

			var result map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if tt.expected.hasError {
				if _, ok := result["type"]; !ok {
					t.Errorf("expected error response to have 'type' field")
				}
			} else {
				if _, ok := result["tasks"]; !ok {
					t.Errorf("expected success response to have 'tasks' field")
				}
			}
		})
	}
}
