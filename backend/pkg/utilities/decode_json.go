package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MalformedRequest struct {
	Status int
	Msg string
}

func (mr *MalformedRequest) Error() string {
	return mr.Msg
}

// DecodeJSONBody is a wrapper around json Decode that sends client-safe and readable errors if any.
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dstn any) error {	
	// If Content-Type header is present, check that it has the value of application/json
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			return &MalformedRequest{
				Status: http.StatusUnsupportedMediaType,
				Msg: msg,
			}
		}
	}

	// Limit the request body to 1MB
	r.Body = http.MaxBytesReader(w, r.Body, 1_048_576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dstn)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var maxBytesError *http.MaxBytesError

		switch {
		// Catch any syntax errors in the JSON and send an error message
        // which interpolates the location of the problem to make it
        // easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg: msg,
			}
        // In some circumstances Decode() may also return an
        // io.ErrUnexpectedEOF error for syntax errors in the JSON. There
        // is an open issue regarding this at
        // https://github.com/golang/go/issues/25956.		
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg: msg,
			}
		// Catch any type errors like trying to assign a string to an int field.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg: msg,
			}
		// Catch the error caused by extra unexpected fields in the request body.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg: msg,
			}
		// io.EOF is returned by Decode() if the request body is empty.
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg: msg,
			}
		// Catch error caused by the request body being too large.
		case errors.As(err, &maxBytesError):
			msg := fmt.Sprintf("Request body must not be larger than %d bytes", maxBytesError.Limit)
			return &MalformedRequest{
				Status: http.StatusRequestEntityTooLarge,
				Msg: msg,
			}
		default:
			return err
		}
	}
    // Call decode again, using a pointer to an empty anonymous struct as 
    // the destination. If the request body only contained a single JSON 
    // object this will return an io.EOF error. So if we get anything else, 
    // we know that there is additional data in the request body.
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		return &MalformedRequest{
			Status: http.StatusBadRequest,
			Msg: msg,
		}
	}

	return nil
}