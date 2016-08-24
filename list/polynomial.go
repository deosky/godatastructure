package list

import "fmt"

type Polyn struct {
	Coef float64
	Expn int
	Next *Polyn
}

//CreatePolyn 输入m项的系数和指数，建立表示一元多项式的有序链表P
func CreatePolyn(p *Polyn, m int) {
	for i := 0; i < m; i++ {
		var coef float64 = 0.00
		var expn int = 0
		fmt.Scanln(&coef, &expn)
		fmt.Printf("Coef = %f , Expn = %d\n", coef, expn)

		node := makeNOde(coef, expn)
		orderInsertPolyn(p, node)
		q := p.Next
		for i := 0; q != nil; i++ {
			fmt.Printf("%f %d -- ", q.Coef, q.Expn)
			q = q.Next
		}
		fmt.Println("\n----------------------------")
	}
}

func makeNOde(coef float64, expn int) *Polyn {
	return &Polyn{Coef: coef, Expn: expn, Next: nil}
}

//OrderInsertPolyn 有序插入多项式节点到链表中, 从小到大
func orderInsertPolyn(p *Polyn, data *Polyn) {
	q := p
	r := p.Next
	for i := 0; r != nil; i++ {

		if compare(data, r) <= 0 {
			data.Next = r
			q.Next = data
			return
		}
		q = r
		r = r.Next
	}
	q.Next = data
}

//compare 比较两个节点的指数大小,如果是如果node1小于node2则返回-1 等于返回0 大于返回1
func compare(node1 *Polyn, node2 *Polyn) int {
	result := node1.Expn - node2.Expn
	if result == 0 {
		return result
	}

	if result > 0 {
		return 1
	}

	return -1
}

//AddPolyn 多项式加法：pa = pa + pb ，利用两个多项式的节点构成和多项式
func AddPolyn(pa, pb *Polyn) {
	pprior := pa
	qa := pa.Next
	qb := pb.Next
	for qa != nil && qb != nil {
		switch compare(qa, qb) {
		case -1:
			pprior = qa
			qa = qa.Next
		case 0:
			if qa.Coef+qb.Coef == 0 {
				//删除节点
				pprior.Next = qa.Next
				qa = pprior.Next
			} else {
				//相加
				qa.Coef = qa.Coef + qb.Coef
				//pprior = qa
				//qa = qa.Next
			}
			qb = qb.Next
		case 1:
			r := qb.Next
			qb.Next = pprior.Next
			pprior.Next = qb

			pprior = qb
			qb = r
		}
	}
	if qb != nil {
		pprior.Next = qb
	}
}
