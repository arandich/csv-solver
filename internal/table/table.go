package table

import (
	"log"
	"regexp"
	"strconv"
)

func FindAndSolve(table [][]string) [][]string {
	rows := make(map[string][]string)
	columns := make(map[int]string)

	for i, v := range table {
		rows[v[0]] = table[i][1:]
	}
	for i, v := range table[0][1:] {
		columns[i] = v
	}
	// Иттерируемся по строчкам из map rows, делаем проверку на пустое значение, чтобы пропустить столбец индексов строк
	for i, v := range rows {
		if i == "" {
			continue
		}
		// После отсеивания 1 столбца идем по массиву ячеек
		for i2, v2 := range v {
			// Проверяем, если первый символ ячейки равен =, то начинаем обрабатывать ячейку
			if string(v2[0]) == "=" {
				// Проверяем правильность написания выражения
				re := regexp.MustCompile(`^=[a-zA-Z]+?[0-9]+[+\*\/\-][a-zA-Z]+?[0-9]+$`)
				ok := re.MatchString(v2)
				if !ok {
					log.Println("Invalid expression")
					return nil
				}
				// Вытаскиваем из строки обозначения столбцов, математический оператор, индекс строчки
				columnNameArr := regexp.MustCompile("[a-zA-Z]+").FindAllString(v2, 2)
				operator := regexp.MustCompile("[+\\*\\/\\-]").FindString(v2)
				rowArr := regexp.MustCompile("[0-9]+").FindAllString(v2, 2)

				var arg1 = ""
				var arg2 = ""
				// Обходим map columns и находим там столбец
				// Записываем в аргумент значение из map rows, в ключ подставляем индекс строчки
				// Так как дальше идет массив, то указываем индекс элемента, отнимаем 1, так как у нас есть столбец для индексов строчки и нам его учитывать не нужно
				// i3-1 = столбец, rowArr[int] = индекс строчки
				for i3, v3 := range columns {
					if columnNameArr[0] == v3 {
						if _, ok = rows[rowArr[0]]; !ok {
							log.Println("Invalid raw number")
							return nil
						}
						arg1 = rows[rowArr[0]][i3]
					}
					if columnNameArr[1] == v3 {
						if _, ok = rows[rowArr[1]]; !ok {
							log.Println("Invalid raw number")
							return nil
						}
						arg2 = rows[rowArr[1]][i3]
					}
				}

				arg1I, err := strconv.Atoi(arg1)
				if err != nil {
					log.Println("Conversion error", err)
					return nil
				}
				arg2I, err := strconv.Atoi(arg2)
				if err != nil {
					log.Println("Conversion error", err)
					return nil
				}
				res := 0
				// В зависимости от полученного оператора производим вычисление
				switch operator {
				case "+":
					res = arg1I + arg2I
				case "-":
					res = arg1I - arg2I
				case "*":
					res = arg1I * arg2I
				case "/":
					if arg1I == 0 {
						log.Println("Impossible to divide 0")
						return nil
					}
					if arg2I == 0 {
						log.Println("Cannot be divided by 0")
						return nil
					}
					res = arg1I / arg2I
				}
				for index, value := range table {
					if value[0] == i {
						table[index][i2+1] = strconv.Itoa(res)
					}
				}
			}
		}
	}
	return table
}
