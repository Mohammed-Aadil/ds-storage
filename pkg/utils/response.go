package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Mohammed-Aadil/ds-storage/pkg/errors"
)

//HTTPHandlerError handler error
type HTTPHandlerError struct {
	Errors         interface{} `json:"errors,omitempty"`
	NonFieldErrors interface{} `json:"nonFieldErrors,omitempty"`
	FieldErrors    interface{} `json:"fieldErrors,omitempty"`
}

//HTTPErrorResponse http error response
func HTTPErrorResponse(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	switch err.(type) {
	case errors.FieldValidationError:
		fieldErr, ok := err.(errors.FieldValidationError)
		if ok {
			json.NewEncoder(w).Encode(&HTTPHandlerError{FieldErrors: &fieldErr})
		}
	case errors.NonFieldValidationError:
		fieldErr, ok := err.(errors.NonFieldValidationError)
		if ok {
			json.NewEncoder(w).Encode(&HTTPHandlerError{NonFieldErrors: &fieldErr})
		}
	default:
		if errType, ok := err.(error); ok && errType.Error() == "EOF" {
			json.NewEncoder(w).Encode(&HTTPHandlerError{Errors: "No data found in request"})
		} else {
			json.NewEncoder(w).Encode(&HTTPHandlerError{Errors: errType.Error()})
		}
	}
}

//HTTPSuccessResponse http success response
func HTTPSuccessResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

//HTTPFileSuccessResponse http file success response
func HTTPFileSuccessResponse(w http.ResponseWriter, r *http.Request, path string) {
	file, err := os.Open(path)
	if err != nil {
		HTTPErrorResponse(w, err)
		return
	}
	defer file.Close()
	fileInfo, errInfo := file.Stat()
	if errInfo != nil {
		HTTPErrorResponse(w, errInfo)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	io.Copy(w, file)
}

//HTTPChannelResponse http channel response
type HTTPChannelResponse struct {
	URL      string
	Body     io.ReadCloser
	Err      error
	FileName string
}
