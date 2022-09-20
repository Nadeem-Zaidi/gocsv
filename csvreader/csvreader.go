package csvreader

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Nadeem-Zaidi/gocsv/errorhandler"
)

func recoverFromPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func ReadCSV(filename string) ([][]string, error) {
	list := make([][]string, 0)

	file := strings.Split(filename, ".")
	defer recoverFromPanic()
	if file[1] == "csv" {

		f, err := os.Open(filename)
		if err != nil {
			panic(err.Error())
		}
		csvReader := csv.NewReader(f)
		for {
			rec, err := csvReader.Read()
			if err == io.EOF {
				break

			}
			if err != nil {
				panic(err.Error())
			}

			list = append(list, rec)

		}
		return list, nil

	} else {
		fmt.Println("file is not csv")
		return nil, &errorhandler.CustomError{
			Statuscode: 999,
			Err:        errors.New("something went wrong in reading the csv file"),
		}

	}

}
