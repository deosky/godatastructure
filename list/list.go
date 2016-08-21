package list

import "errors"

const (
	LIST_INT_SIZE int = 10
	LISTINCREMENT int = 10
	ERROR         int = -1
)

type ElemType int

//Status 1是成功， 0 是失败
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

//DeleteSqList 删除指定位置的元素(pos从1开始)
func DeleteSqList(l *SqList, pos int) (ElemType, error) {
	if pos < 1 || pos > l.Length {
		return 0, errors.New("无效的删除位置")
	}

	var elem ElemType
	elem = l.Elem[pos-1]

	for i := pos; i < l.Length; i++ {
		l.Elem[i-1] = l.Elem[i]
	}
	l.Length--

	return elem, nil
}

//LocateSqElem 在l中找到第一个满足compare()元素的位序,如果不存在则返回0(位置从1开始)
func LocateSqElem(l *SqList, e ElemType, compare func(e1, e2 ElemType) Status) int {
	for i := 0; i < l.Length; i++ {
		if compare(e, l.Elem[i]) == 1 {
			return i + 1
		}
	}
	return 0
}
