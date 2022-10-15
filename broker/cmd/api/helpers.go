package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrJsonRead error = errors.New("[JSON read error]")

func ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	var maxBytes int64 = 1048576 // 1 MB

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return fmt.Errorf("%w: body must have only a single JSON value", ErrJsonRead)
	}

	return nil
}

func WriteJson(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)

	return err
}

func ErrorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := JsonResponse{
		Error:   true,
		Message: err.Error(),
	}

	return WriteJson(w, statusCode, payload)
}
