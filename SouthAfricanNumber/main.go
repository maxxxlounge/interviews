package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/handler"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	var fileSource = flag.String("i", "input.csv", "-i input source file")
	var storeFile = flag.String("d", "output.json", "-d destination file (json)")
	var port = flag.String("p", "8888", "-p listen port")
	flag.Parse()

	l := log.New()
	reader, err := os.Open(*fileSource)
	DieOnErr(err)
	defer reader.Close()

	validNumbers := make(map[string]*NumberManager.Row)
	loadedNumbers := make(map[string]*NumberManager.Row)
	criticalNumbers := make(map[string]*NumberManager.Row)
	fixableNumbers := make(map[string]*NumberManager.Row)

	l.Info("loading csv file")
	r := csv.NewReader(reader)
	//remove header
	_,err = r.Read()
	if err != nil {
		l.Fatal(err)
	}
	rowindex := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		DieOnErr(err)

		//prevent missing columns
		if len(record) < 2 {
			DieOnErr(errors.Errorf("bad input file format, missing column at line %v", rowindex))
		}

		// prevent duplicated index
		if _, ok := loadedNumbers[record[0]]; ok {
			log.Fatalf("duplicated index '%v' on row '%v' ", record[0], rowindex)
		}

		row := NumberManager.New(record[1])
		loadedNumbers[record[0]] = row
		rowindex++
	}
	l.Info("processing numbers...")

	for k, v := range loadedNumbers {
		switch v.Type {
		case NumberManager.ValidFirstAttempt:
			validNumbers[k] = v
		case NumberManager.InvalidCritical:
			criticalNumbers[k] = v
		case NumberManager.InvalidButFixable:
			fixableNumbers[k] = v
		}
	}

	log.Infof	("given numbers %v\n", len(loadedNumbers))
	log.Infof("---")
	log.Infof("valid numbers %v\n", len(validNumbers))
	log.Infof("Fixable numbers %v\n", len(fixableNumbers))
	log.Infof("Critical numbers %v\n", len(criticalNumbers))
	log.Infof("---")
	log.Infof("Counter Sum  %v\n", len(criticalNumbers)+len(fixableNumbers)+len(validNumbers))

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
		w.Header().Set("Access-Control-Allow-Origin:", "*")
		w.Header().Add("Content-Type", "application/json")
		p1 := r.URL.Query().Get("number")
		handler.Check(w, p1)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	s := &http.Server{
		Addr:           ":" + *port,
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	l.Infof("saving data on '%s ", *storeFile)
	err = Store(loadedNumbers, *storeFile)
	if err != nil {
		// not die because the service still up
		err = errors.Wrap(err, "error storing data on file")
		log.Warn(err)
	}
	l.Infof("starting endpoint on port :%s", *port)
	log.Fatal(s.ListenAndServe())
}

func Store(m map[string]*NumberManager.Row, filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	out, err := json.Marshal(m)
	if err != nil {
		return err
	}
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

func DieOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
