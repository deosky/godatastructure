package main

import "godatastructure/list"
import "fmt"
import "strconv"

func main() {
	listNode := list.SqList{}
	list.InitSqList(&listNode)
	for i := 0; i < 15; i++ {
		list.InsertSqList(&listNode, i+1, list.ElemType(i))
		fmt.Println("len = " + strconv.Itoa(listNode.Length) + " cap = " + strconv.Itoa(listNode.Listsize))
	}
	//fmt.Println("abc")
	fmt.Println(listNode.Length)
}
