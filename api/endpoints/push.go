package endpoints

import (
	"encoding/json"
	"net/http"

	helpers "github.com/dwburke/go-tools/gorillamuxhelpers"
	"github.com/gorilla/mux"
	//"github.com/spf13/viper"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/push", Push).Methods("POST")
}

type Event struct {
	Name  string
	Value string
	Type  string
}

func Push(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var e Event
	if err := decoder.Decode(&e); err != nil {
		helpers.RespondWithError(w, 500, err.Error())
		return
	}

	if e.Name == "" {
		helpers.RespondWithError(w, 500, "required param 'name' is missing")
		return
	}
	if e.Value == "" {
		helpers.RespondWithError(w, 500, "required param 'value' is missing")
		return
	}
	if e.Type == "" {
		helpers.RespondWithError(w, 500, "required param 'type' is missing")
		return
	}

	helpers.RespondWithJSON(w, 200, map[string]interface{}{
		"ok":    1,
		"value": e.Value,
		"type":  e.Type,
		"name":  e.Name,
	})
}
