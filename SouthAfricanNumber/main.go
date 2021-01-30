package main

import (
	"encoding/csv"
	"github.com/maxxxlounge/interviews/SouthAfricanNumber/NumberManager"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/errors/fmt"
	"io"
	"os"
)

func main() {

	loadedNumbers := make(map[string]*NumberManager.Row)
	criticalNumbers := make(map[string]*NumberManager.Row)
	fixableNumbers := make(map[string]*NumberManager.Row)


	if len(os.Args) < 2 {
		log.Fatal("missing input file in args")
	}

	filepath := os.Args[1]
	reader, err := os.Open(filepath)
	defer reader.Close()
	DieOnErr(err)
	r := csv.NewReader(reader)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		DieOnErr(err)
		row, err := NumberManager.New(record[1])
		DieOnErr(err)
		loadedNumbers[record[0]] = row
	}

	for k, v := range loadedNumbers {
		errStr := ""
		err := NumberManager.FindCriticalError(v.GetOriginalGivenNumber())
		if err != nil {
			criticalNumbers[k] = v
			errStr = err.Error()
		}else{
			fixableNumbers[k] = v
		}
		fmt.Printf("%v\t%v\t%v\n", k, v.GetOriginalGivenNumber(), errStr)
	}

	fmt.Println()
	fmt.Printf("given numbers %v\n",len(loadedNumbers))
	fmt.Printf("Critical numbers %v\n",len(criticalNumbers))
	fmt.Printf("Fixable numbers %v\n",len(fixableNumbers))
}

func DieOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
