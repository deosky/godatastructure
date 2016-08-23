package list

import "fmt"

type Polyn struct {
	coef float64
	expn int
	next *Polyn
}

//CreatePolyn 输入m项的系数和指数，建立表示一元多项式的有序链表P
func CreatePolyn(p *Polyn, m int) {
	for i := 0; i < m; i++ {
		var coef float64 = 0.00
		var expn int = 0
		fmt.Scanln(&coef, &expn)
		fmt.Printf("coef = %f , expn = %d\n", coef, expn)

		node := makeNOde(coef, expn)
		orderInsertPolyn(p, node)
		q := p.next
		for i := 0; q != nil; i++ {
			fmt.Printf("%f %d -- ", q.coef, q.expn)
			q = q.next
		}
		fmt.Println("\n----------------------------")
	}
}

func makeNOde(coef float64, expn int) *Polyn {
	return &Polyn{coef: coef, expn: expn, next: nil}
}

//OrderInsertPolyn 有序插入多项式节点到链表中, 从小到大
func orderInsertPolyn(p *Polyn, data *Polyn) {
	q := p
	r := p.next
	for i := 0; r != nil; i++ {

		if compare(data, r) <= 0 {
			data.next = r
			q.next = data
			return
		}
		q = r
		r = r.next
	}
	q.next = data
}

//compare 比较两个节点的指数大小,如果是如果node1小于node2则返回-1 等于返回0 大于返回1
func compare(node1 *Polyn, node2 *Polyn) int {
	return node1.expn - node2.expn
}
