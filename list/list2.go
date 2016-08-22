package list

import "errors"
import "fmt"

//LNode 链表节点
type LNode struct {
	Data ElemType
	Next *LNode
}

//GetElem_L 返回 l 中第i个数据元素的值
//l为带头节点的单链表的头指针
//i 从1开始
//算法 2.8
func GetElem_L(l *LNode, i int) (ElemType, error) {
	p := l.Next
	j := 1

	for p != nil && j < i {
		p = p.Next
		j++
	}

	if p == nil || j > i {
		return 0, errors.New("第i个元素不存在")
	}

	return p.Data, nil
}

//ListInsert_L 在第i个元素之前插入节点
//l为带头节点的单链表的头指针
//i 从1开始
//算法 2.9
func ListInsert_L(l *LNode, i int, elem ElemType) error {

	if i < 1 {
		return errors.New("无效的链表位置")
	}
	p := l
	j := 0
	for ; p != nil && j < i-1; j++ {
		p = p.Next
	}
	if j > i || p == nil {
		return errors.New("i小于1或者大于表长加1")
	}
	node := &LNode{Data: elem, Next: p.Next}
	p.Next = node

	return nil
}

//ListDelete_L 删除位置i处的节点
//l为带头节点的单链表的头指针
//i 从1开始
//算法 2.10
func ListDelete_L(l *LNode, i int) (ElemType, error) {
	if i < 1 {
		return 0, errors.New("删除的位置不合理")
	}
	p := l
	j := 0
	for ; p != nil && j < i-1; j++ {
		p = p.Next
	}

	if p == nil || p.Next == nil || j > i {
		return 0, errors.New("删除的位置不合理")
	}

	q := p.Next
	p.Next = p.Next.Next

	return q.Data, nil
}

//CreateList_L 从尾到表头逆向建立单链表
//算法2.11
func CreateList_L(l *LNode, n int) {
	for i := 0; i < n; i++ {
		p := &LNode{}
		num := 0
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		p.Data = ElemType(num)
		p.Next = l.Next
		l.Next = p
	}
}

//MergeList_L 一直单链表la和lb的元素按值非递减排序
//归并la和lb 得到新的单链线性表lc ,lc的元素也按值非递减排序
//算法2.12
func MergeList_L(la *LNode, lb *LNode) *LNode {
	pa, pb := la.Next, lb.Next

	pchead := &LNode{}
	pc := pchead
	for pa != nil && pb != nil {
		if pa.Data > pb.Data {
			pc.Next = pb
			pb = pb.Next
		} else {
			pc.Next = pa
			pa = pa.Next
		}
		pc = pc.Next
	}

	if pa != nil {
		pc.Next = pa
	} else {
		pc.Next = pb
	}
	return pchead
}
