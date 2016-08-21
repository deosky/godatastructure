package list

import "errors"

const (
	LIST_INT_SIZE int = 10
	LISTINCREMENT int = 10
	ERROR         int = -1
)

type ElemType int
type Status int

type SqList struct {
	Elem     []ElemType
	Length   int
	Listsize int
}

func InitSqList(l *SqList) Status {
	array := [LIST_INT_SIZE]ElemType{}
	l.Elem = array[:]
	l.Length = 0
	l.Listsize = LIST_INT_SIZE

	return 1
}

//InsertSqList 在pos(从1开始)之前插入一个元素
func InsertSqList(l *SqList, pos int, elem ElemType) (Status, error) {
	if pos < 1 || pos > l.Length+1 {
		return 0, errors.New("无效的插入位置")
	}
	if l.Length >= l.Listsize {
		incrementArray := [LISTINCREMENT]ElemType{}
		l.Elem = append(l.Elem, incrementArray[:]...)
		l.Listsize = cap(l.Elem)
	}
	for i := l.Length - 1; i >= pos-1; i-- {
		l.Elem[i+1] = l.Elem[i]
	}
	l.Elem[pos-1] = elem
	l.Length++

	return 1, nil
}
