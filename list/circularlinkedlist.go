package list //循环链表
import "errors"

//DuLNode 双向循环链接
type DuLNode struct {
	Data  ElemType
	prior *DuLNode
	next  *DuLNode
}

//ListInsert_DuL 在带头节点的双循环线性表L中第i个位置之前插入元素e i的合法值为 1<= i<= 表长 + 1
//算法 2.18
func ListInsert_DuL(l *DuLNode, i int, e ElemType) error {
	//h := l
	p := l
	for j := 0; p != nil && j <= i; j++ {
		p = p.next
	}

	if p == nil {
		return errors.New("无效的双链循环线性表L")
	}

	s := &DuLNode{Data: e}
	s.prior = p.prior
	p.prior.next = s
	s.next = p
	p.prior = s

	return nil
}

//ListDelete_DuL 删除带头节点的双链循环线性表L的第i个元素 ， i 的合法值为 1<= i <=表长
//算法 2.19
func ListDelete_DuL(l *DuLNode, i int) (ElemType, error) {
	p := l
	for j := 0; p != nil && j < i; j++ {
		p = p.next
	}
	if p == nil {
		return 0, errors.New("无效的双链循环线性表L")
	}

	p.prior.next = p.next
	p.next.prior = p.prior

	return p.Data, nil
}
