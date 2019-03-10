package endpoints_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"

	"github.com/dwburke/stdata/api/endpoints"
)

func setup() (r *mux.Router) {
	r = mux.NewRouter()

	r.HandleFunc("/push", endpoints.Push).Methods("POST")

	return
}

func TestPush(t *testing.T) {
	r := setup()

	params := map[string]interface{}{
		"name":  "silly test name",
		"value": "off",
		"type":  "whatami",
	}
	jsonstr, err := json.Marshal(params)
	if err != nil {
		t.Errorf("Error marshalling params to json: %s", err)
	}

	req, _ := http.NewRequest("POST", "/push", bytes.NewReader([]byte(jsonstr)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expect(t, w.Code, http.StatusOK, w)

	body, err := ioutil.ReadAll(w.Body)
	expect(t, err, nil, "")
	expect(t, w.Code, http.StatusOK, string(body))

	type Result struct {
		Ok    int
		Name  string
		Value string
		Type  string
	}

	obj := &Result{}

	err = json.Unmarshal(body, &obj)
	expect(t, err, nil, "")

	expect(t, obj.Value, "off", "value")
	expect(t, obj.Name, "silly test name", "name")
	expect(t, obj.Type, "whatami", "type")
	expect(t, obj.Ok, 1, "ok")

}

func expect(t *testing.T, a interface{}, b interface{}, body interface{}) {
	if a != b {
		if body == "" {
			t.Errorf("Expected [%v] (type %v) - Got [%v] (type %v)",
				b, reflect.TypeOf(b), a, reflect.TypeOf(a))
		} else {
			t.Errorf("Expected [%v] (type %v) - Got [%v] (type %v) : %#v",
				b, reflect.TypeOf(b), a, reflect.TypeOf(a), body)
		}
	}
}
