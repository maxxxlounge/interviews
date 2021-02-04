package handler

import (
	"encoding/json"
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func Check(w http.ResponseWriter, number string) {
	if number == "" {
		http.Error(w, "missing number", http.StatusBadRequest)
		return
	}
	row := NumberManager.New(number)
	out, err := json.Marshal(row)
	if err != nil {
		err = errors.Wrap(err, "error marshalling number information to JSON")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if row.Type == NumberManager.ValidFirstAttempt {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write(out)
	if err != nil {
		log.Printf("error on write: %v", err)
	}
}

func ShowNumbers(w http.ResponseWriter, numberlist map[string]*NumberManager.Row) {
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
