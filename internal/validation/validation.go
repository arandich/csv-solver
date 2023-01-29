package validation

import (
	"log"
	"os"
	"regexp"
)

type Validation interface {
	ValidateName(string) bool
	ValidateTable(*os.File, rune) [][]string
}

type Validate struct {
	Validation
}

func (v Validate) ValidateName(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z1-9_]+?\.csv`)
	ok := re.MatchString(s)
	if !ok {
		return ok
	}
	return ok
}

func (v Validate) ValidateTable(table [][]string) bool {

	if table[0][0] != "" {
		log.Println("Invalid value of the first cell")
		return false
	}

	re := regexp.MustCompile(`^[a-zA-Z]+?[0-9]*`)
	for i, value := range table[0][1:] {
		ok := re.MatchString(value)
		if !ok {
			log.Println("Invalid column name value: ", i+1, "Value: ", value)
			return false
		}
	}

	re = regexp.MustCompile(`^[0-9]+`)
	for i, value := range table[1:] {
		ok := re.MatchString(value[0])
		if !ok {
			log.Println("Incorrect line numbering: ", i, "Value: ", value)
			return false
		}
		for i2, v2 := range value {
			if v2 == "" {
				log.Println("Invalid cell value: ", v2, "row: ", i+1, "column: ", i2+1)
				return false
			} else {
				re = regexp.MustCompile(`^=[a-zA-Z]+?[0-9]+[+\*\/\-][a-zA-Z]+?[0-9]+$`)
				match1 := re.MatchString(v2)
				re = regexp.MustCompile(`^-?[0-9]+`)
				match2 := re.MatchString(v2)
				if match1 || match2 {

				} else {
					log.Println("Invalid cell value: ", v2, "row: ", i+1, "column: ", i2+1)
					return false
				}
			}
		}
	}

	return true
}
