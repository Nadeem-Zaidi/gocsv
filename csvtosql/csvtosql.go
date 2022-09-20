package csvtosql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Nadeem-Zaidi/gocsv/csvreader"
	_ "github.com/go-sql-driver/mysql"
)

func CscvToSql(filename string) {
	result, err := csvreader.ReadCSV(filename)
	if err != nil {
		log.Fatal(err)
	}

	outerValuesList := make([]string, 0)
	valuesList := make([]string, 0)
	for oi, oe := range result {
		for ii := range oe {
			if result[oi][ii] == "null" || result[oi][ii] == "" || result[oi][ii] == " " {
				result[oi][ii] = "NULL"
				valuesList = append(valuesList, result[oi][ii])

			} else {

				r := fmt.Sprintf("'%s'", result[oi][ii])
				valuesList = append(valuesList, r)

			}

		}
		vr := fmt.Sprintf("(%s)", strings.Join(valuesList, ","))
		outerValuesList = append(outerValuesList, vr)
		valuesList = nil

	}

	headerString := strings.ReplaceAll(outerValuesList[0], "'", "`")
	fmt.Println(headerString)

	fmt.Println(strings.Join(outerValuesList[1:], ","))

	db, err := sql.Open("mysql", "root:owl@tcp(127.0.0.1:3306)/fastapi")
	if err != nil {
		fmt.Println(err)
	}
	query := fmt.Sprintf("INSERT INTO category %s VALUES %s", headerString, strings.Join(outerValuesList[1:], ","))
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	lastid, err := res.LastInsertId()
	if err != nil {
		fmt.Println("there")
		fmt.Println(err.Error())
	}

	fmt.Printf("The last inserted row id: %d\n", lastid)

}
