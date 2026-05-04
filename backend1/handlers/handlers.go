package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/calculator/backend/calculator"
)

// CalculateRequest represents the JSON body for a calculation.
type CalculateRequest struct {
	Operation string   `json:"operation"`
	A         *float64 `json:"a"`
	B         *float64 `json:"b,omitempty"`
}

// CalculateResponse is returned to the client.
type CalculateResponse struct {
	Result float64 `json:"result"`
}

// ErrorResponse is returned on failure.
type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// CalculateHandler handles POST /api/calculate
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "method not allowed"})
		return
	}

	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid JSON body"})
		return
	}

	if req.A == nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "field 'a' is required"})
		return
	}

	// Operations that only need one operand
	switch req.Operation {
	case "sqrt":
		result, err := calculator.SquareRoot(*req.A)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, CalculateResponse{Result: result})
		return
	}

	// All other operations require two operands
	if req.B == nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "field 'b' is required for this operation"})
		return
	}

	a, b := *req.A, *req.B

	var result float64
	switch req.Operation {
	case "add":
		result = calculator.Add(a, b)
	case "subtract":
		result = calculator.Subtract(a, b)
	case "multiply":
		result = calculator.Multiply(a, b)
	case "divide":
		res, err := calculator.Divide(a, b)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}
		result = res
	case "exponentiate":
		result = calculator.Exponentiate(a, b)
	case "percentage":
		result = calculator.Percentage(a, b)
	default:
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "unsupported operation: " + req.Operation})
		return
	}

	writeJSON(w, http.StatusOK, CalculateResponse{Result: result})
}

// HealthHandler handles GET /api/health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
