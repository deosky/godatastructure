package list

const (
	LIST_INT_SIZE int = 100
	LISTINCREMENT int = 10
)

type ElemType int
type Status int

type SqList struct {
	Elem     []ElemType
	Length   int
	listsize int
}

func InitList(l *SqList) Status {
	array := [LIST_INT_SIZE]ElemType{}
	l.Elem = array[:]
	l.Length = 0
	l.listsize = LIST_INT_SIZE

	return 1
}
