package main

import (
	"encoding/csv"
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/handler"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/errors/fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("missing input file in args")
	}

	l := log.New()

	filepath := os.Args[1]
	reader, err := os.Open(filepath)
	defer reader.Close()
	DieOnErr(err)

	validNumbers := make(map[string]*NumberManager.Row)
	loadedNumbers := make(map[string]*NumberManager.Row)
	criticalNumbers := make(map[string]*NumberManager.Row)
	fixableNumbers := make(map[string]*NumberManager.Row)

	l.Info("loading csv file")
	r := csv.NewReader(reader)
	//remove header
	r.Read()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		DieOnErr(err)
		row := NumberManager.New(record[1])
		loadedNumbers[record[0]] = row
	}
	l.Info("processing numbers...")

	for k, v := range loadedNumbers {
		switch v.Type {
		case NumberManager.ValidFirstAttempt:
			validNumbers[k] = v
			break
		case NumberManager.InvalidCritical:
			criticalNumbers[k] = v
			break
		case NumberManager.InvalidButFixable:
			fixableNumbers[k] = v
			break
		}
	}

	fmt.Println()
	fmt.Printf("given numbers %v\n", len(loadedNumbers))
	fmt.Println("---")
	fmt.Printf("valid numbers %v\n", len(validNumbers))
	fmt.Printf("Fixable numbers %v\n", len(fixableNumbers))
	fmt.Printf("Critical numbers %v\n", len(criticalNumbers))
	fmt.Println("---")
	fmt.Printf("Counter Sum  %v\n", len(criticalNumbers)+len(fixableNumbers)+len(validNumbers))

	l.Info("starting endpoing on port")
	var h http.Handler
	http.HandleFunc("/numbers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		p := r.URL.Query().Get("type")
		switch p {
		case "valid":
			handler.ShowNumbers(w, validNumbers)
		case "critical":
			handler.ShowNumbers(w, criticalNumbers)
		case "fixable":
			handler.ShowNumbers(w, fixableNumbers)
		default:
			handler.ShowNumbers(w, loadedNumbers)
		}
	})
	http.HandleFunc("/numbers/check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		p1 := r.URL.Query().Get("number")
		handler.Check(w, p1)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/search.html")
	})

	s := &http.Server{
		Addr:           ":8888",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func Print(m map[string]NumberManager.Row) {
	for k, v := range m {
		var errOutput string
		for ei, e := range v.Errors {
			if ei > 0 {
				errOutput += ", "
			}
			errOutput += e
		}
		out := fmt.Sprintf("%v\t%v\t%v:\t%v", k, v.Original, v.Type, errOutput)
		if v.Type == NumberManager.InvalidButFixable {
			out += " " + v.Original + "\t"
		}
		fmt.Println(out)
	}
}

func DieOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
