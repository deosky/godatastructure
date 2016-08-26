package stack

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	//op1 行 opt2在列
	//+ - * / ( ) # ,大于为> ，小于为<,等于为= ,x为非法的运算
	//op1\opt2  +  -  *  /  (  )  #
	//       +  >  >  <  <  <  >  >
	//       -
	//       *
	//       /
	//       (  <  <  <  <  <  =  x
	//       )
	//       #
	precedeTable = [][]byte{
		{'>', '>', '<', '<', '<', '>', '>'},
		{'>', '>', '<', '<', '<', '>', '>'},
		{'>', '>', '>', '>', '<', '>', '>'},
		{'>', '>', '>', '>', '<', '>', '>'},
		{'<', '<', '<', '<', '<', '=', 'x'},
		{'>', '>', '>', '>', ' ', '>', '>'},
		{'<', '<', '<', '<', '<', 'x', '='},
	}
)

//EvaluateExpression 表达式求值 ,只可以计算整数
func EvaluateExpression() error {
	//操作符
	optr := &SqStack{}
	optr.Push(int('#'))
	//操作数
	opnd := &SqStack{}
	var str string
	str = "(3*(7-2))-((3*2)-7)#"
	var num []byte
	for i := 0; i < len(str); i++ {

		if isDigit(str[i]) {
			num = append(num, str[i])
		} else if isOperator(str[i]) {
			if num != nil {
				d, _ := strconv.Atoi(string(num))
				num = nil
				opnd.Push(d)
			}
			op1, _ := optr.Peek()
			switch Precede(byte(op1), str[i]) {
			case '>':
				num2, _ := opnd.Pop()
				num1, _ := opnd.Pop()
				opc, _ := optr.Pop()
				result, _ := calc(num1, num2, byte(opc))
				opnd.Push(result)
				//继续比较当前位置的符号,是否小于前一个符号
				i--

			case '<':
				optr.Push(int(str[i]))
			case '=':
				optr.Pop()
			case 'x':
				return errors.New("无效的表达式")
			}

		} else {
			return errors.New("无效的表达式")
		}
	}
	e, _ := opnd.Peek()
	fmt.Printf("e = %d \n", e)
	return nil
}

func calc(num1, num2 int, op byte) (int, error) {
	switch op {
	case '+':
		return num1 + num2, nil
	case '-':
		return num1 - num2, nil
	case '*':
		return num1 * num2, nil
	case '/':
		return num1 / num2, nil
	}
	return 0, errors.New("无效的操作符")
}

func isDigit(digit byte) bool {
	return digit >= '0' && digit <= '9'
}

func isOperator(operator byte) bool {
	switch operator {
	case '+':
		fallthrough
	case '-':
		fallthrough
	case '*':
		fallthrough
	case '/':
		fallthrough
	case '(':
		fallthrough
	case ')':
		fallthrough
	case '#':
		return true
	}
	return false
}

//Precede 比较两个运算符的优先级
func Precede(op1, op2 byte) byte {
	return precedeTable[getOpIndex(op1)][getOpIndex(op2)]
}

//getOpIndex 获取指定运算符的位置
func getOpIndex(op byte) int {
	ops := []byte{'+', '-', '*', '/', '(', ')', '#'}
	for i, v := range ops {
		if v == op {
			return i
		}
	}
	return -1
}
