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

	if len(os.Args) < 2 {
		log.Fatal("missing input file in args")
	}

	filepath := os.Args[1]
	reader, err := os.Open(filepath)
	defer reader.Close()
	DieOnErr(err)

	validNumbers := make(map[string]*NumberManager.Row)
	loadedNumbers := make(map[string]*NumberManager.Row)
	criticalNumbers := make(map[string]*NumberManager.Row)
	fixableNumbers := make(map[string]*NumberManager.Row)

	r := csv.NewReader(reader)
	//remove header
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
		isValidAtFirstAttempt := NumberManager.IsRightFormat(v.GetOriginalNumber())
		if isValidAtFirstAttempt {
			validNumbers[k] = v
			validNumbers[k].Type = NumberManager.ValidFirstAttempt
			continue
		}
		err := NumberManager.FindCriticalError(v.GetOriginalNumber())
		if err != nil {
			criticalNumbers[k] = v
			criticalNumbers[k].Type = NumberManager.InvalidCritical
			criticalNumbers[k].Errors = append(criticalNumbers[k].Errors, err)
			continue
		}
		fixableNumbers[k] = v
		fixableNumbers[k].Type = NumberManager.InvalidButFixable
	}

	for k, v := range loadedNumbers {
		var errOutput string
		for ei, e := range v.Errors {
			if ei > 0 {
				errOutput += ", "
			}
			errOutput += e.Error()
		}
		out := fmt.Sprintf("%v\t%v\t%v:\t%v", k, v.GetOriginalNumber(), v.Type, errOutput)
		if v.Type == NumberManager.InvalidButFixable {
			out += " " + v.GetChangedNumber() + "\t"
		}
		fmt.Println(out)
	}

	fmt.Println()
	fmt.Printf("given numbers %v\n", len(loadedNumbers))
	fmt.Printf("valid numbers %v\n", len(validNumbers))
	fmt.Printf("Fixable numbers %v\n", len(fixableNumbers))
	fmt.Printf("Critical numbers %v\n", len(criticalNumbers))
}

func DieOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
