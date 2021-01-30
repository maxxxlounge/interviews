package main

import (
	"encoding/csv"
	"io"
	"os"
	log "github.com/sirupsen/logrus"
)

type row struct{
	originalNumber string
	ChangedNumber string
	Err error
}

var numbers map[string]row

func main(){

	if len(os.Args) < 2{
		log.Fatal("missing input file in args")
	}

	filepath := os.Args[1]
	reader,err  := os.Open(filepath)
	DieOnErr(err)
	r := csv.NewReader(reader)
	for {
		record,err := r.Read()
		if err == io.EOF {
			break
		}
		DieOnErr(err)
		row,err := GenerateRow(record[1])
		DieOnErr(err)
		numbers[record[0]] = row
	}
}

func GenerateRow(number string)(row,error){
	r := row{
		originalNumber: number,
	}
	correctedNr, err := CorrectNumber(number)
	r.ChangedNumber = correctedNr
	r.Err = err
	return r,nil
}

func CorrectNumber(number string)(string, error){
 panic("not implemented")
}


func DieOnErr(err error){
	if err!=nil{
		log.Fatal(err)
	}
}
