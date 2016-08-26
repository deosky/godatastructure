package stack

/*
#include<stdio.h>
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	MAXSIZE int = 100
)

// var (
// 	stackArray [MAXSIZE]int
// )

type SqStack struct {
	base       int
	top        int
	StackSize  int
	stackArray [MAXSIZE]int
}

//Init 初始化一个栈
func (s *SqStack) Init() {
	s.base = 0
	s.top = 0
	s.StackSize = MAXSIZE
}

//Push 推入栈中一个元素
func (s *SqStack) Push(e int) error {
	if s.top-s.base >= MAXSIZE {
		return errors.New("栈满")
	}

	s.stackArray[s.top] = e
	s.top++

	return nil
}

//Pop 出栈一个元素
func (s *SqStack) Pop() (int, error) {
	if s.top <= s.base {
		return 0, errors.New("栈空")
	}

	e := s.stackArray[s.top-1]
	s.top--

	return e, nil
}

//Peek 查看一下栈顶元素的值
func (s *SqStack) Peek() (int, error) {
	if s.top <= s.base {
		return 0, errors.New("栈空")
	}

	return s.stackArray[s.top-1], nil
}

//Conversion 数制转换
func Conversion(num int) {
	stack := &SqStack{}
	stack.Init()
	n := num
	for n != 0 {
		stack.Push(n % 8)
		n /= 8
	}

	for {
		e, err := stack.Pop()
		if err != nil {
			break
		}
		fmt.Printf("%d ", e)
	}
	fmt.Println()
}

//getParenthesis 获取配对的字符
func getParenthesis(c int) int {
	switch c {
	case int('('):
		return int(')')
	case int('['):
		return int(']')
	case int('{'):
		return int('}')
	}
	return 0
}

//ParenthesisMatching 括号匹配
func ParenthesisMatching(data string) bool {
	if len(data)%2 != 0 || len(data) <= 0 {
		return false
	}

	stack := &SqStack{}
	stack.Init()

	for i := 0; i < len(data); i++ {
		switch data[i] {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			stack.Push(int(data[i]))
		default:
			e, err := stack.Pop()
			if err != nil {
				return false
			}
			if getParenthesis(e) != int(data[i]) {
				return false
			}
		}
	}

	return true
}

//LineEdit 行编辑程序
func LineEdit() {
	reader := bufio.NewReader(os.Stdin)
	stack := &SqStack{}
	stack.Init()
	for {
		c, _ := reader.ReadByte()
		if c == '\r' || c == '\n' {
			break
		}
		switch c {
		case '#':
			stack.Pop()
		case '@':
			for {
				_, err := stack.Pop()
				if err != nil {
					return
				}
			}
		default:

			stack.Push(int(c))
		}
	}

	stack1 := &SqStack{}
	stack1.Init()
	for {

		e, err := stack.Pop()

		if err != nil {
			break
		}

		stack1.Push(e)
	}
	for {
		e, err := stack1.Pop()
		if err != nil {
			break
		}
		fmt.Printf("%c", e)
	}
}
