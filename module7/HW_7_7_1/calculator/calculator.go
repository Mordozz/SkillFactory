package calc

import (
	"fmt"

)

type calculator struct{
	number1 float64
	number2 float64
	operator string
}


func NewCalculator(number1, number2 float64, operator string) *calculator {
	return &calculator{
		number1: number1,
		number2: number2,
		operator: operator,
	}
}

func (c *calculator) Calculate() (float64, error){
	switch c.operator {
    case "+":
        return c.add(c.number1, c.number2)
    case "-":
        return c.sub(c.number1, c.number2)
    case "*":
        return c.mul(c.number1, c.number2)
    case "/":
        return c.div(c.number1, c.number2)
    default:
        return 0, fmt.Errorf("неизвестный оператор")
    }
}
	

func (c *calculator) add(a, b float64) (float64, error) {
    return a + b, nil
}

func (c *calculator) sub(a, b float64) (float64, error) {
    return a - b, nil
}

func (c *calculator) mul(a, b float64) (float64, error) {
    return a * b, nil
}

func (c *calculator) div(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль запрещено")
    }
    return a / b, nil
}