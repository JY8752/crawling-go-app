package json

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func JsonDecode[T comparable](r *http.Request, w http.ResponseWriter, params *T) *T {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return params
}

func JsonEncode[T comparable](buff *bytes.Buffer, v T) error {
	enc := json.NewEncoder(buff)
	return enc.Encode(v)
}
