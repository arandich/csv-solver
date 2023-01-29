package tests

import (
	"csv_solver/internal/table"
	"csv_solver/internal/validation"
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"testing"
)

type ValidationNameTest struct {
	arg1     string
	expected bool
}

type ValidationTableTest struct {
	name     string
	expected bool
}

type FormatAndSolveTest struct {
	table    [][]string
	expected [][]string
}

var ValidationTableTests = []ValidationTableTest{
	{"true_test.csv", true},
	{"invalid_test1.csv", false},
	{"invalid_test2.csv", false},
	{"invalid_test3.csv", false},
	{"invalid_test4.csv", false},
}

var ValidationNameTests = []ValidationNameTest{
	{"test.csv", true},
	{"test.exe", false},
	{".test.csv", false},
	{"test.234csv", false},
}

var FormatAndSolveTests = []FormatAndSolveTest{
	{table: [][]string{
		{"", "A", "B", "Cell"},
		{"1", "1", "0", "1"},
		{"2", "2", "=A1+Cell30", "0"},
		{"30", "0", "=B1+A1", "5"},
	}, expected: [][]string{
		{"", "A", "B", "Cell"},
		{"1", "1", "0", "1"},
		{"2", "2", "6", "0"},
		{"30", "0", "1", "5"},
	}},
	{table: [][]string{
		{"", "A", "B", "Cell", "F", "G"},
		{"1", "1", "0", "1", "=A1*A30", "6"},
		{"2", "2", "=A1+Cell30", "0", "=B1*A30", "3"},
		{"30", "0", "=B1+A1", "5", "=G30/A2", "8"},
	}, expected: [][]string{
		{"", "A", "B", "Cell", "F", "G"},
		{"1", "1", "0", "1", "0", "6"},
		{"2", "2", "6", "0", "0", "3"},
		{"30", "0", "1", "5", "4", "8"},
	}},
	{table: [][]string{
		{"", "A", "B", "Cell", "F", "G"},
		{"1", "1", "-", "1", "=A1*A30", "6"},
		{"2", "2", "=A1+Cell30", "0", "=B1*A30", "3"},
		{"30", "0", "=B1+A1", "5", "=G30/A2", "8"},
	}, expected: nil},
}

func TestValidationName(t *testing.T) {
	validate := validation.Validate{}
	for _, test := range ValidationNameTests {
		if output := validate.ValidateName(test.arg1); output != test.expected {
			t.Errorf("Test: %s Output %t not equal to expected %t", test.arg1, output, test.expected)
		}
	}
}

func TestValidationTable(t *testing.T) {
	validate := validation.Validate{}
	for _, test := range ValidationTableTests {
		file, err := os.Open("./csv/" + test.name)
		if err != nil {
			log.Println(err)
			return
		}

		reader := csv.NewReader(file)
		reader.Comma = ','

		tableCsv, err := reader.ReadAll()
		if err != nil {
			log.Println("Некорректная таблица")
			return
		}
		file.Close()

		if output := validate.ValidateTable(tableCsv); output != test.expected {
			t.Errorf("Test: %s Output %t not equal to expected %t", test.name, output, test.expected)
		}
	}
}

func TestFromatAndSolve(t *testing.T) {
	for _, test := range FormatAndSolveTests {
		if output := table.FindAndSolve(test.table); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
