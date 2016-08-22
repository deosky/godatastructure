package list

import "fmt"

const (
	MAXSIZE = 100
)

type SLNode struct {
	Data ElemType
	Cur  int
}

type SLinkList [MAXSIZE]SLNode

//LocateElem_SL 在静态单链线性表中查找第一个值为e的元素
//若找到，则返回它在S中的位序，否则返回0
//算法2.13
func LocateElem_SL(s *SLinkList, e ElemType) int {

	i := s[0].Cur
	for i != 0 && s[i].Data != e {
		i = s[i].Cur
	}

	return i
}

//InitSpace_SL 将一维数组space中各分量链成一个备用链表 ， space[0].Cur 为头指针
//算法 2.14
func InitSpace_SL(space *SLinkList) {
	for i := 0; i < MAXSIZE-1; i++ {
		space[i].Cur = i + 1
	}
	space[MAXSIZE-1].Cur = 0
}

//Malloc_SL 若备用空间链表非空,则返回分配的节点下标 ，否则返回0
//算法 2.15
func Malloc_SL(space *SLinkList) int {
	i := space[0].Cur
	if i != 0 {
		space[0].Cur = space[i].Cur
	}
	return i
}

//Free_SL 把下标为k的空闲节点回收到备用链表
//算法 2.16
func Free_SL(space *SLinkList, k int) {
	space[k].Cur = space[0].Cur
	space[0].Cur = k
}

//Difference 依次蔬菜集合A和B的元素,在一维数组space中建立表示集合（A-B）并（B-A）
//的静态链表，S为其头指针，假设备用空间足够大，space[0].Cur为其头指针
//算法  2.17
func Difference(space *SLinkList, s *int) {
	InitSpace_SL(space)
	*s = Malloc_SL(space)
	r := *s
	m, n := 0, 0
	fmt.Scanln(&m, &n)
	fmt.Printf("m = %d , n = %d\r\n", m, n)
	for ; m > 0; m-- {
		var num1 int = 0
		fmt.Scanf("%c\r\n", &num1)
		fmt.Printf("mum = %c \r\n", num1)
		i := Malloc_SL(space)
		space[i].Data = ElemType(num1)
		space[i].Cur = space[r].Cur
		space[r].Cur = i
		r = i
	}

	space[r].Cur = 0
	h := space[*s].Cur
	for h != 0 {
		fmt.Printf("%c ", space[h].Data)
		h = space[h].Cur
	}

	fmt.Println("\n----------------")

	for ; n > 0; n-- {
		var num2 int = 0
		var k int
		fmt.Scanf("%c\r\n", &num2)
		fmt.Printf("mum = %c \r\n", num2)
		p := *s
		k = space[*s].Cur
		for k != space[r].Cur && space[k].Data != ElemType(num2) {
			p = k
			k = space[k].Cur
		}
		if k == space[r].Cur {
			i := Malloc_SL(space)
			space[i].Data = ElemType(num2)
			space[i].Cur = space[r].Cur
			space[r].Cur = i
		} else {
			space[p].Cur = space[k].Cur
			Free_SL(space, k)
			if r == k {
				r = p
			}
		}

		h := space[*s].Cur
		for h != 0 {
			fmt.Printf("%c ", space[h].Data)
			h = space[h].Cur
		}
		fmt.Println("\n----------------")
	}

}
