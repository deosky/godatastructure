package stack

import "errors"
import "fmt"

const (
	MAXSIZE int = 100
)

var (
	stackArray [MAXSIZE]int
)

type SqStack struct {
	base      int
	top       int
	StackSize int
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

	stackArray[s.top] = e
	s.top++

	return nil
}

//Pop 出栈一个元素
func (s *SqStack) Pop() (int, error) {
	if s.top <= s.base {
		return 0, errors.New("栈空")
	}

	e := stackArray[s.top-1]
	s.top--

	return e, nil
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
