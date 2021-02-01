package handler

import (
	"encoding/json"
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/pkg/errors"
	"net/http"
)

func ShowNumbers(w http.ResponseWriter, numberlist map[string]*NumberManager.Row) {
	out, err := json.Marshal(numberlist)
	if err != nil {
		err = errors.Wrap(err, "error marshalling list to JSON")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(out)
}
