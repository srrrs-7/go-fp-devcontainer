package tasks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestPutHandler(t *testing.T) {
	type args struct {
		formData map[string]string
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
				formData: map[string]string{
					"id":          "550e8400-e29b-41d4-a716-446655440000",
					"title":       "Updated Task",
					"description": "Updated Description",
					"completed":   "true",
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
				formData: map[string]string{
					"id":    "invalid-uuid",
					"title": "Updated Task",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				hasError:   true,
			},
		},
		{
			testName: "missing id",
			args: args{
				formData: map[string]string{
					"title": "Updated Task",
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
				formData: map[string]string{
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
			formData := url.Values{}
			for k, v := range tt.args.formData {
				formData.Set(k, v)
			}
			req := httptest.NewRequest(http.MethodPut, "/tasks", strings.NewReader(formData.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()
			PutHandler(w, req)

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
				if _, ok := result["id"]; !ok {
					t.Errorf("expected success response to have 'id' field")
				}
			}
		})
	}
}
