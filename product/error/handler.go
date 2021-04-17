package error

import (
	"encoding/json"
	"fmt"
	"time"
)

type HTTPError struct {
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Err       string    `json:"error"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
}

func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

func NewHTTPError(status int, message string, errs string, path string) *HTTPError {
	return &HTTPError{
		Timestamp: time.Now(),
		Message:   message,
		Status:    status,
		Err:       errs,
		Path:      path,
	}
}
