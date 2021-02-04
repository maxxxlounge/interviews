package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maxxxlounge/interviews/SouthAfricanNumber/numbermanager"
	"github.com/pkg/errors"
)

func Check(w http.ResponseWriter, number string) {
	if number == "" {
		http.Error(w, "missing number", http.StatusBadRequest)
		return
	}
	row := numbermanager.New(number)
	out, err := json.Marshal(row)
	if err != nil {
		err = errors.Wrap(err, "error marshalling number information to JSON")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if row.Type == numbermanager.ValidFirstAttempt {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write(out)
	if err != nil {
		log.Printf("error on write: %v", err)
	}
}

func ShowNumbers(w http.ResponseWriter, numberlist map[string]*numbermanager.Row) {
	out, err := json.Marshal(numberlist)
	if err != nil {
		err = errors.Wrap(err, "error marshalling list to JSON")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(out)
	if err != nil {
		log.Printf("error on write: %v", err)
	}
}
