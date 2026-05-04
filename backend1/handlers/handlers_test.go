package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func postCalculate(t *testing.T, body interface{}) *httptest.ResponseRecorder {
	t.Helper()
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal request: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	CalculateHandler(rr, req)
	return rr
}

func float64Ptr(v float64) *float64 { return &v }

func TestCalculateHandler_Add(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "add", A: float64Ptr(2), B: float64Ptr(3)})
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d; want 200", rr.Code)
	}
	var resp CalculateResponse
	json.NewDecoder(rr.Body).Decode(&resp)
	if resp.Result != 5 {
		t.Errorf("result = %v; want 5", resp.Result)
	}
}

func TestCalculateHandler_Subtract(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "subtract", A: float64Ptr(10), B: float64Ptr(4)})
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d; want 200", rr.Code)
	}
	var resp CalculateResponse
	json.NewDecoder(rr.Body).Decode(&resp)
	if resp.Result != 6 {
		t.Errorf("result = %v; want 6", resp.Result)
	}
}

func TestCalculateHandler_Multiply(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "multiply", A: float64Ptr(3), B: float64Ptr(7)})
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d; want 200", rr.Code)
	}
	var resp CalculateResponse
	json.NewDecoder(rr.Body).Decode(&resp)
	if resp.Result != 21 {
		t.Errorf("result = %v; want 21", resp.Result)
	}
}

func TestCalculateHandler_Divide(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "divide", A: float64Ptr(10), B: float64Ptr(4)})
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d; want 200", rr.Code)
	}
	var resp CalculateResponse
	json.NewDecoder(rr.Body).Decode(&resp)
	if resp.Result != 2.5 {
		t.Errorf("result = %v; want 2.5", resp.Result)
	}
}

func TestCalculateHandler_DivideByZero(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "divide", A: float64Ptr(10), B: float64Ptr(0)})
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d; want 400", rr.Code)
	}
	var resp ErrorResponse
	json.NewDecoder(rr.Body).Decode(&resp)
	if resp.Error != "division by zero" {
		t.Errorf("error = %q; want 'division by zero'", resp.Error)
	}
}

func TestCalculateHandler_Sqrt(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "sqrt", A: float64Ptr(16)})
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d; want 200", rr.Code)
	}
	var resp CalculateResponse
	json.NewDecoder(rr.Body).Decode(&resp)
	if resp.Result != 4 {
		t.Errorf("result = %v; want 4", resp.Result)
	}
}

func TestCalculateHandler_SqrtNegative(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "sqrt", A: float64Ptr(-4)})
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d; want 400", rr.Code)
	}
}

func TestCalculateHandler_Exponentiate(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "exponentiate", A: float64Ptr(2), B: float64Ptr(10)})
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d; want 200", rr.Code)
	}
	var resp CalculateResponse
	json.NewDecoder(rr.Body).Decode(&resp)
	if resp.Result != 1024 {
		t.Errorf("result = %v; want 1024", resp.Result)
	}
}

func TestCalculateHandler_Percentage(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "percentage", A: float64Ptr(200), B: float64Ptr(15)})
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d; want 200", rr.Code)
	}
	var resp CalculateResponse
	json.NewDecoder(rr.Body).Decode(&resp)
	if resp.Result != 30 {
		t.Errorf("result = %v; want 30", resp.Result)
	}
}

func TestCalculateHandler_UnsupportedOperation(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "modulo", A: float64Ptr(5), B: float64Ptr(2)})
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d; want 400", rr.Code)
	}
}

func TestCalculateHandler_MissingA(t *testing.T) {
	rr := postCalculate(t, map[string]interface{}{"operation": "add", "b": 5})
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d; want 400", rr.Code)
	}
}

func TestCalculateHandler_MissingB(t *testing.T) {
	rr := postCalculate(t, CalculateRequest{Operation: "add", A: float64Ptr(5)})
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d; want 400", rr.Code)
	}
}

func TestCalculateHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/calculate", bytes.NewReader([]byte("not json")))
	rr := httptest.NewRecorder()
	CalculateHandler(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d; want 400", rr.Code)
	}
}

func TestCalculateHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/calculate", nil)
	rr := httptest.NewRecorder()
	CalculateHandler(rr, req)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d; want 405", rr.Code)
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rr := httptest.NewRecorder()
	HealthHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d; want 200", rr.Code)
	}
}
