package response

import (
	"api/src/domain/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testResponse struct {
	Message string `json:"message,omitempty"`
	ID      string `json:"id,omitempty"`
}

func TestOK(t *testing.T) {
	type args struct {
		body any
	}
	type expected struct {
		statusCode int
		body       map[string]string
	}

	tests := []struct {
		testName string
		args     args
		expected expected
	}{
		{
			testName: "success",
			args: args{
				body: testResponse{Message: "success"},
			},
			expected: expected{
				statusCode: http.StatusOK,
				body:       map[string]string{"message": "success"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			w := httptest.NewRecorder()
			OK(w, tt.args.body)

			resp := w.Result()
			if resp.StatusCode != tt.expected.statusCode {
				t.Errorf("expected status %v, got %v", tt.expected.statusCode, resp.Status)
			}

			if resp.Header.Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type application/json, got %v", resp.Header.Get("Content-Type"))
			}

			var result map[string]string
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			for k, v := range tt.expected.body {
				if result[k] != v {
					t.Errorf("expected %s %v, got %v", k, v, result[k])
				}
			}
		})
	}
}

func TestCreated(t *testing.T) {
	type args struct {
		body any
	}
	type expected struct {
		statusCode int
		body       map[string]string
	}

	tests := []struct {
		testName string
		args     args
		expected expected
	}{
		{
			testName: "created",
			args: args{
				body: testResponse{ID: "123"},
			},
			expected: expected{
				statusCode: http.StatusCreated,
				body:       map[string]string{"id": "123"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			w := httptest.NewRecorder()
			Created(w, tt.args.body)

			resp := w.Result()
			if resp.StatusCode != tt.expected.statusCode {
				t.Errorf("expected status %v, got %v", tt.expected.statusCode, resp.Status)
			}

			if resp.Header.Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type application/json, got %v", resp.Header.Get("Content-Type"))
			}

			var result map[string]string
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			for k, v := range tt.expected.body {
				if result[k] != v {
					t.Errorf("expected %s %v, got %v", k, v, result[k])
				}
			}
		})
	}
}

func TestHandleAppError(t *testing.T) {
	type args struct {
		err model.AppError
	}
	type expected struct {
		statusCode int
		body       map[string]string
	}

	tests := []struct {
		testName string
		args     args
		expected expected
	}{
		{
			testName: "bad request",
			args: args{
				err: model.NewBadRequestError(nil, "TestDomain"),
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: map[string]string{
					"type":   model.BadRequestErrorName,
					"domain": "TestDomain",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			w := httptest.NewRecorder()
			HandleAppError(w, tt.args.err)

			resp := w.Result()
			if resp.StatusCode != tt.expected.statusCode {
				t.Errorf("expected status %v, got %v", tt.expected.statusCode, resp.Status)
			}

			if resp.Header.Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type application/json, got %v", resp.Header.Get("Content-Type"))
			}

			var result map[string]string
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			for k, v := range tt.expected.body {
				if result[k] != v {
					t.Errorf("expected %s %v, got %v", k, v, result[k])
				}
			}
		})
	}
}
