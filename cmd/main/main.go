package main

import (
	"bufio"
	"csv_solver/internal/table"
	"csv_solver/internal/validation"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Write file name Example: test.csv\nThe file must be located in the root directory of the project")
	var fileName string
	_, err := fmt.Scan(&fileName)
	if err != nil {
		log.Println(err)
		return
	}

	validate := validation.Validate{}
	ok := validate.ValidateName(fileName)
	if !ok {
		log.Println("Invalid file name")
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return
	}

	reader := csv.NewReader(file)
	reader.Comma = ','

	tableCsv, err := reader.ReadAll()
	if err != nil {
		log.Println("Incorrect table")
		return
	}

	file.Close()

	ok = validate.ValidateTable(tableCsv)
	if !ok {
		log.Println("Invalid table")
		return
	}

	tableCsv = table.FindAndSolve(tableCsv)
	if tableCsv == nil {
		return
	}

	w := csv.NewWriter(os.Stdout)

	err = w.WriteAll(tableCsv)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Press 'q' to quit")
		}
	}

}
