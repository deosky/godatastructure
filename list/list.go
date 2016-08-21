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

//InitSqList 处理是线性表
//算法2.3
func InitSqList(l *SqList) Status {
	array := [LIST_INT_SIZE]ElemType{}
	l.Elem = array[:]
	l.Length = 0
	l.Listsize = LIST_INT_SIZE

	return 1
}

//InsertSqList 在pos(从1开始)之前插入一个元素
//算法2.4
func InsertSqList(l *SqList, pos int, elem ElemType) (Status, error) {
	if pos < 1 || pos > l.Length+1 {
		return 0, errors.New("无效的插入位置")
	}
	if l.Length >= l.Listsize {
		incrementArray := [LISTINCREMENT]ElemType{}
		slice := append(l.Elem, incrementArray[:]...)
		l.Elem = slice[:cap(slice)]
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
//算法2.5
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
//算法2.6
func LocateSqElem(l *SqList, e ElemType, compare func(e1, e2 ElemType) Status) int {
	for i := 0; i < l.Length; i++ {
		if compare(e, l.Elem[i]) == 1 {
			return i + 1
		}
	}
	return 0
}

//MergeSqList 顺序表合并
//算法2.7
func MergeSqList(la, lb SqList) *SqList {

	var lc SqList
	InitSqList(&lc)
	var i int
	var j int
	for i < la.Length && j < lb.Length {
		if la.Elem[i] > lb.Elem[j] {
			InsertSqList(&lc, lc.Length+1, lb.Elem[j])
			j++
		} else {
			InsertSqList(&lc, lc.Length+1, la.Elem[i])
			i++
		}
	}

	for i < la.Length {
		InsertSqList(&lc, lc.Length+1, la.Elem[i])
		i++
	}
	for j < lb.Length {
		InsertSqList(&lc, lc.Length+1, lb.Elem[j])
		j++
	}

	return &lc
}

//MergeSqList2 顺序表合并
//算法2.2
func MergeSqList2(la, lb SqList) *SqList {
	return MergeSqList(la, lb)
}

//UnionSqList 将所有在线性表lb中但不在线性表la中的数据插入到la中
//算法2.1
func UnionSqList(la, lb *SqList) {

	for i := 0; i < lb.Length; i++ {
		if LocateSqElem(la, lb.Elem[i], func(e1, e2 ElemType) Status {
			if e1 != e2 {
				return Status(0)
			}

			return Status(1)
		}) < 1 {
			InsertSqList(la, la.Length+1, lb.Elem[i])
		}
	}
}
