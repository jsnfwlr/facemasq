package httperror

import (
	"database/sql"
	"facemasq/lib/logging"
	"facemasq/lib/translate"
	"io"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type HTTPError struct {
	statusCode int
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(err error) HTTPError {
	switch err {
	case io.EOF:
		return HTTPError{
			statusCode: http.StatusBadRequest,
			Message:    "EOF reading HTTP request body",
		}
	case sql.ErrNoRows:
		return HTTPError{
			statusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		}
	}

	return HTTPError{
		statusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}
}

func Handler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(out http.ResponseWriter, in bunrouter.Request) error {
		// Call the next handler on the chain to get the error.
		err := next(out, in)
		logging.Error("error with request: %v", err)
		switch err := err.(type) {
		case nil:
			// no error
		case HTTPError: // already a HTTPError
			err.translate()
			out.WriteHeader(err.statusCode)
			_ = bunrouter.JSON(out, err)
		default:
			httpErr := NewHTTPError(err)
			httpErr.translate()
			out.WriteHeader(httpErr.statusCode)
			_ = bunrouter.JSON(out, httpErr)
		}

		return err // return the err in case there other middlewares
	}
}

func (err *HTTPError) translate() {
	err.Message = translate.Message(err.Code, err.Message)
}
