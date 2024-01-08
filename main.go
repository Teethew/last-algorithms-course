package main

import (
	"container/list"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	examples := make(map[int][]string)

	examples[9] = []string{"2", "1", "+", "3", "*"}
	examples[6] = []string{"4", "13", "5", "/", "+"}
	examples[22] = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	for result, expression := range examples {
		fmt.Printf("result: %v | evalRPN(): %v\n", result, evalRPN(expression))
	}
}

func evalRPN(tokens []string) int {

	stack := list.New()

	for _, token := range tokens {
		if !isOperation(token) {
			stack.PushFront(token)
			continue
		}

		operand1, _ := strconv.Atoi(stack.Front().Value.(string))

		operand2, _ := strconv.Atoi(stack.Front().Next().Value.(string))

		var result int

		switch token {
		case "+":
			result = operand1 + operand2
		case "-":
			result = operand1 - operand2
		case "/":
			result = operand1 / operand2
		case "*":
			result = operand1 * operand2
		}

		stack.Remove(stack.Front())
		stack.Front().Value = fmt.Sprint(result)
	}

	result, _ := strconv.Atoi(stack.Front().Value.(string))

	return result
}

func isOperation(s string) bool {
	operations := []string{"/", "+", "*", "-"}
	return slices.Contains[[]string, string](operations, s)
}
