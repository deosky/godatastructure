package main

import "godatastructure/list"
import "fmt"

func main() {
	listNode := list.SqList{}
	list.InitList(&listNode)
	for i := 0; i < 100; i++ {
		listNode.Elem[i] = list.ElemType(i)
		listNode.Length++
	}
	fmt.Println("abc")
	fmt.Println(listNode)
}
