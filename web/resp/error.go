package resp

import (
	"net/http"
)

const requestID = "X-Request-ID"

// Error represents a response object of api error
type Error struct {
	ID         string                 `json:"id,omitempty"`
	Message    string                 `json:"message,omitempty"`
	Details    map[string]interface{} `json:"details,omitempty"`
	StackTrace string                 `json:"stackTrace,omitempty"`
}

// ServeBadRequest serves http BadRequest
func ServeBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	re := &JSONResponse{
		response: response{code: http.StatusBadRequest},
		Errors: []Error{
			{
				ID:      getRequestID(r),
				Message: err.Error(),
			},
		},
	}
	Render(w, r, re)
}

// ServeNotFound serves http NotFound
func ServeNotFound(w http.ResponseWriter, r *http.Request, err error) {
	re := &JSONResponse{
		response: response{code: http.StatusNotFound},
		Errors: []Error{
			{
				ID:      getRequestID(r),
				Message: err.Error(),
			},
		},
	}
	Render(w, r, re)
}

// ServeUnauthorized serves http Unauthorized
func ServeUnauthorized(w http.ResponseWriter, r *http.Request, err error) {
	re := &JSONResponse{
		response: response{code: http.StatusUnauthorized},
		Errors: []Error{
			{
				ID:      getRequestID(r),
				Message: err.Error(),
			},
		},
	}
	Render(w, r, re)
}

// ServeForbidden serves http Forbidden
func ServeForbidden(w http.ResponseWriter, r *http.Request, err error) {
	re := &JSONResponse{
		response: response{code: http.StatusForbidden},
		Errors: []Error{
			{
				ID:      getRequestID(r),
				Message: err.Error(),
			},
		},
	}
	Render(w, r, re)
}

// ServeUnprocessableEntity serves http UnprocessableEntity
func ServeUnprocessableEntity(w http.ResponseWriter, r *http.Request, err error, dtl map[string]interface{}) {
	re := &JSONResponse{
		response: response{code: http.StatusUnprocessableEntity},
		Errors: []Error{
			{
				ID:      getRequestID(r),
				Message: err.Error(),
				Details: dtl,
			},
		},
	}
	Render(w, r, re)
}

// ServeInternalServerError serves http InternalServerError
func ServeInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	re := &JSONResponse{
		response: response{code: http.StatusInternalServerError},
		Errors: []Error{
			{
				ID:      getRequestID(r),
				Message: err.Error(),
			},
		},
	}
	Render(w, r, re)
}

// ServeTooManyRequestsError serves http TooManyRequests
func ServeTooManyRequestsError(w http.ResponseWriter, r *http.Request, err error) {
	re := &JSONResponse{
		response: response{code: http.StatusTooManyRequests},
		Errors: []Error{
			{
				ID:      getRequestID(r),
				Message: err.Error(),
			},
		},
	}
	Render(w, r, re)
}

// ServeError serves error with appropriate http status code determined from error type
func ServeError(w http.ResponseWriter, r *http.Request, err error) {
	ServeInternalServerError(w, r, err)
}

func getRequestID(r *http.Request) string {
	return r.Header.Get(requestID)
}
