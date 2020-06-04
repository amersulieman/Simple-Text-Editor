package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Stack to keep track of memory
type Stack []string

//Push addes to the stack
func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

//Pop  pops from the stack and return value popped
func (s *Stack) Pop() string {
	*s = (*s)[:len(*s)-1]
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]

}

//Op1 is Appending to the old string
func op1(oldString, newString string) string {
	return fmt.Sprintf("%v%v", oldString, newString)
}

//Op2 is deleting the last k chars
func op2(value string, numOfChar int) string {
	charachterToStopAt := len(value) - numOfChar
	newValue := value[:charachterToStopAt]
	return newValue
}

func main() {

	var stack Stack
	S := ""

	reader := bufio.NewReader(os.Stdin)
	numOps, _ := reader.ReadString('\n')
	op := strings.Fields(numOps)
	operationsNumber, _ := strconv.Atoi(op[0])

	//Loop for number of operations given
	for i := 0; i < operationsNumber; i++ {
		input, _ := reader.ReadString('\n')
		op := strings.Fields(input)

		switch option := op[0]; option {
		case "1":
			S = op1(S, op[1])
			stack.Push(S)

		case "2":
			numChars, _ := strconv.Atoi(op[1])
			S = op2(S, numChars)
			stack.Push(S)
		case "3":
			k, _ := strconv.Atoi(op[1])
			fmt.Println(string(S[k-1]))
		case "4":
			S = stack.Pop()
		}
	}
}
