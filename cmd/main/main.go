package main

import (
	"bufio"
	"csv_solver/internal/table"
	"csv_solver/internal/validation"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

var fileArg = flag.String("file", "", "file-name")

func main() {
	flag.Parse()
	var fileName string
	if *fileArg == "" {
		fmt.Println("Write file name Example: test.csv\nThe file must be located in the root directory of the project")
		_, err := fmt.Scan(&fileName)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		fileName = *fileArg
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
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','

	tableCsv, err := reader.ReadAll()
	if err != nil {
		log.Println("Incorrect table")
		return
	}

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
