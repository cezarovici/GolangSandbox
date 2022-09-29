package main

import "log"

// "A string representing an expression is given. The expression contains
// numbers and arithmetic operations (only =E2=80=98+=E2=80=99 and =E2=80=98=
// -=E2=80=99). Write code that
// calculates the value of an expression. It is forbidden to use ready-made
// parsers created for such a purpose. Example:
// "2+6+4" -12
// "20-10-6" -4
// "123+200" -323"

type expression struct {
	itSelf   string
	values   []int
	operands []byte
}

func (e *expression) setItSelf(itSelf string) {
	e.itSelf = itSelf
}

func (e *expression) setOperands() {
	for _, element := range e.itSelf {
		if isOperand(byte(element)) {
			e.operands = append(e.operands, byte(element))
		}
	}
}

func (e *expression) setValues() {
	var left int

	for i, element := range e.itSelf {
		if isOperand(byte(element)) {
			e.values = append(e.values, stringToNumber(e.itSelf[left:i]))
			left = i + 1
			log.Print(e.values)
		}
	}
}

func (e *expression) setInfos(exp string) {
	e.setItSelf(exp)
	e.setOperands()
	e.setValues()
}

func (e expression) printExpresion() {
	log.Println("From ", e.itSelf)
	log.Println("Have:\n-", e.operands, "\n-", e.values)
}

func isOperand(character byte) bool {
	if character == '+' || character == '-' || character == '/' || character == '*' || character == '%' {
		return true
	}
	return false
}

func stringToNumber(value string) int {
	var res int

	for _, element := range value {
		res = res*10 + (int(element) - 48)
	}

	return res
}

const exp = "2+6+4"

func main() {
	// var e expression

	// e.setInfos(exp)
	// e.printExpresion()

	log.Print(stringToNumber("abc"))
}
