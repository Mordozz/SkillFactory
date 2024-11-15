package main

import (
    "bufio"
    "fmt"
    "calc/calculator" 
    "os"
    "strconv"
    "strings"
)

func isOperatorValid(operator string) bool{
	switch operator{
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}
func main() {

	for{
		var operator string
    	reader := bufio.NewReader(os.Stdin)

		fmt.Print("Введите первое число: ")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)
		
        for {
            fmt.Print("Введите оператор (+, -, *, /): ")
            opInput, err := reader.ReadString('\n')
            if err != nil {
                fmt.Println("Ошибка при чтении оператора:", err)
                continue
            }
            operator = strings.TrimSpace(opInput)
            if isOperatorValid(operator) {
                break
            } else {
                fmt.Println("Недопустимый оператор. Пожалуйста, введите один из следующих операторов: +, -, *, /")
            }
        }


		fmt.Print("Введите второе число: ")
		input2, _ := reader.ReadString('\n')
		input2 = strings.TrimSpace(input2)

		number1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("Ошибка конвертации числа 1:", err)
			return
		}

		number2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Ошибка конвертации числа 2:", err)
			return
		}

		calcInitiate := calc.NewCalculator(number1, number2, operator)
		result, err := calcInitiate.Calculate()
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		fmt.Printf("Результат: %.2f\n", result)
	}
}


