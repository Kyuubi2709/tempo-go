package client

import "fmt"

// JSONRPCRequest represents a JSON-RPC 2.0 request.
type JSONRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      interface{}   `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

// JSONRPCResponse represents a JSON-RPC 2.0 response.
type JSONRPCResponse struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      interface{}   `json:"id"`
	Result  interface{}   `json:"result,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

// JSONRPCError represents a JSON-RPC 2.0 error.
type JSONRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error implements the error interface for JSONRPCError.
func (e *JSONRPCError) Error() string {
	return fmt.Sprintf("RPC error %d: %s", e.Code, e.Message)
}

// CheckError returns an error if the response contains an RPC error.
func (r *JSONRPCResponse) CheckError() error {
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// JSON-RPC error codes (standard + custom)
const (
	// Standard JSON-RPC errors
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603

	// Custom errors for fee payer server
	InvalidTransactionType = -32000
)

// NewJSONRPCRequest creates a new JSON-RPC request.
func NewJSONRPCRequest(id interface{}, method string, params ...interface{}) *JSONRPCRequest {
	return &JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}
}

// NewJSONRPCResponse creates a successful JSON-RPC response.
func NewJSONRPCResponse(id interface{}, result interface{}) *JSONRPCResponse {
	return &JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}
}

// NewJSONRPCErrorResponse creates an error JSON-RPC response.
func NewJSONRPCErrorResponse(id interface{}, code int, message string, data interface{}) *JSONRPCResponse {
	return &JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &JSONRPCError{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
}

// BatchRequest represents a batch of JSON-RPC requests.
// Use NewBatchRequest() to create a new batch and Add() to add requests.
type BatchRequest struct {
	requests []*JSONRPCRequest
	nextID   int
}

// NewBatchRequest creates a new batch request.
func NewBatchRequest() *BatchRequest {
	return &BatchRequest{
		requests: make([]*JSONRPCRequest, 0),
		nextID:   1,
	}
}

// Add adds a request to the batch.
// The ID is automatically assigned incrementally.
func (b *BatchRequest) Add(method string, params ...interface{}) *BatchRequest {
	b.requests = append(b.requests, NewJSONRPCRequest(b.nextID, method, params...))
	b.nextID++
	return b
}

// Requests returns the list of requests in the batch.
func (b *BatchRequest) Requests() []*JSONRPCRequest {
	return b.requests
}

// Len returns the number of requests in the batch.
func (b *BatchRequest) Len() int {
	return len(b.requests)
}
