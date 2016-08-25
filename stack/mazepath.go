package stack

import (
	"errors"
	"fmt"
)

type mazeStack struct {
	data [100]PosType
	base int
	top  int
}

func (m *mazeStack) Push(point PosType) error {
	if m.top-m.base >= 100 {
		return errors.New("栈满")
	}
	m.data[m.top] = point
	m.top++
	return nil
}

func (m *mazeStack) Pop() (PosType, error) {
	if m.top <= m.base {
		return PosType{}, errors.New("栈空")

	}
	e := m.data[m.top-1]
	m.top--
	return e, nil

}

func (m *mazeStack) IsEmpty() bool {
	if m.top <= m.base {
		return true
	}

	return false

}

type PosType struct {
	X   int
	Y   int
	Dir int
}

//MazePath 迷宫求解 X Y 代表行和列
func MazePath(maze [][]int, start, end PosType) {

	mazestack := mazeStack{}

	// // maze = [][]int{
	// // 	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	// // 	{1, 0, 0, 1, 0, 0, 0, 1, 0, 1},
	// // 	{1, 0, 0, 1, 0, 0, 0, 1, 0, 1},
	// // 	{1, 0, 0, 0, 0, 1, 1, 0, 0, 1},
	// // 	{1, 0, 1, 1, 1, 0, 0, 0, 0, 1},
	// // 	{1, 0, 0, 0, 1, 0, 0, 0, 0, 1},
	// // 	{1, 0, 1, 0, 0, 0, 1, 0, 0, 1},
	// // 	{1, 0, 1, 1, 1, 0, 1, 1, 0, 1},
	// // 	{1, 1, 0, 0, 0, 0, 0, 0, 0, 1},
	// // 	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	// // }
	// maze = [][]int{
	// 	{1, 1, 1, 1, 1, 1, 1},
	// 	{1, 0, 1, 0, 0, 0, 1},
	// 	{1, 0, 1, 1, 1, 0, 1},
	// 	{1, 0, 0, 0, 0, 0, 1},
	// 	{1, 0, 1, 1, 1, 0, 1},
	// 	{1, 0, 0, 0, 1, 0, 1},
	// 	{1, 1, 1, 1, 1, 1, 1},
	// }

	// //设置迷宫的起点
	// start = PosType{X: 1, Y: 1, Dir: 1}
	// //设置迷宫的终点
	// end = PosType{X: 5, Y: 5, Dir: 1}

	//将迷宫的启动设置为第一个判断节点
	curpos := start

	for ok := true; ok || !mazestack.IsEmpty(); ok = false {
		//判断当前位置是否通行
		if isPass(maze, curpos) {
			//留下痕迹
			footPrint(maze, curpos)
			//将位置入栈
			if err := mazestack.Push(curpos); err != nil {
				fmt.Println("迷宫路径太长")
				return
			}

			//判断当前位置是否是迷宫的终点
			if curpos.X == end.X && curpos.Y == end.Y {
				//找到了一条通路
				for i := 0; i < len(maze); i++ {
					for j := 0; j < len(maze[0]); j++ {
						fmt.Printf("%d ", maze[i][j])
					}
					fmt.Println()
				}
				for !mazestack.IsEmpty() {
					data, _ := mazestack.Pop()
					fmt.Printf("(%02d %02d)->", data.X, data.Y)
				}
				fmt.Println()
				return
			}
			curpos.Dir = 1
			curpos = nextPos(maze, curpos)

		} else { //当前位置不通
			pos, _ := mazestack.Pop()
			//判断不通节点的上一个节点的方向是否是4，如果是则在该位置已不可能找到通路
			for pos.Dir == 4 {
				//标记当前节点一定不通
				markPrint(maze, pos)
				//直到找到可通的节点
				pos, _ = mazestack.Pop()
			}
			if pos.Dir < 4 {
				pos.Dir++
				//将刚才出栈的有效节点重新入栈
				mazestack.Push(pos)
				//获取下一个位置的节点判断是否
				curpos = nextPos(maze, pos)
			}
		}
	}
}

//isPass 判断当前位置是否可以通行
func isPass(maze [][]int, pos PosType) bool {
	if maze[pos.X][pos.Y] == 0 {
		return true
	}

	return false
}

//nextPos 获取当前节点指向方向的下一个节点 东南西北(1,2,3,4)
func nextPos(maze [][]int, pos PosType) PosType {
	switch pos.Dir {
	case 1:
		return PosType{X: pos.X, Y: pos.Y + 1, Dir: 1}
	case 2:
		return PosType{X: pos.X + 1, Y: pos.Y, Dir: 1}
	case 3:
		return PosType{X: pos.X, Y: pos.Y - 1, Dir: 1}
	case 4:
		return PosType{X: pos.X - 1, Y: pos.Y, Dir: 1}
	}

	return PosType{}
}

//markPrint 标记当前迷宫节点一定不通
func markPrint(maze [][]int, pos PosType) {
	maze[pos.X][pos.Y] = 2
}

//footPrint 留下已经访问的足迹
func footPrint(maze [][]int, pos PosType) {
	maze[pos.X][pos.Y] = 3
}
