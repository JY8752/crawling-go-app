package json

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	//given
	reqBody := bytes.NewBufferString(`{"value": "Hello World"}`)
	r := httptest.NewRequest(http.MethodGet, "https://test.com", reqBody)
	w := httptest.NewRecorder()

	var params struct {
		Value string `json:"value"`
	}

	//when
	result := *JsonDecode(r, w, &params)

	//then
	ex := struct {
		Value string `json:"value"`
	}{Value: "Hello World"}
	if !reflect.DeepEqual(result, ex) {
		t.Errorf("expected %v, act %v\n", ex, result)
	}
}

func TestJsonEncode(t *testing.T) {
	//given
	buff := &bytes.Buffer{}
	v := struct {
		Value string `json:"value"`
	}{Value: "Hello World"}

	//when
	if err := JsonEncode(buff, v); err != nil {
		t.Errorf("failed encode.")
	}

	//then
	ex := `{"value":"Hello World"}`
	act := strings.TrimSuffix(buff.String(), "\n")
	if act != ex {
		t.Errorf("expect: %v, act: %v\n", ex, act)
	}
}
