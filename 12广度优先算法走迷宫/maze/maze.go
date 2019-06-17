package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

/*
	采用广度优先算法实现走迷宫

	一个6行5列的迷宫 只能上下左右移动，0：代表陆地可移动，1：代表墙不可移动
	如下所示：
					| 0 1 2 3 4
				____|___________
				  0 | 0 1 0 0 0
				  1 | 0 0 0 1 0
				  2 | 0 1 0 1 0
				  3 | 1 1 1 0 0
				  4 | 0 1 0 0 1
				  5 | 0 1 0 0 0

*/

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//判断 Point 是否在迷宫内。并返回该 point 在迷宫中的值
func (p point) isInMazeAndGetPointValue(grid [][]int) (int, bool) {

	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

/*
	获取当前执行文件的目录
*/
func currentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(" Can not get current file info")
	}
	lastIndex := strings.LastIndex(file, "/") + 1
	file = file[:lastIndex]
	return file
}

// 首先读取迷宫资源
func readMaze(fileName string) [][]int {
	fmt.Println(fileName)
	file, e := os.Open(fileName)
	if e != nil {
		panic(e)
	}

	var row, col int
	// 从文件中按照给出的格式将值读出写到变量中
	_, e = fmt.Fscanf(file, "%d %d", &row, &col)
	if e != nil {
		panic(e)
	}

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

// 假设 {0,0} 为基准点 Y定义上左下右四个点
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

// 走迷宫
func walkMaze(maze [][]int, start point, end point) [][]int {
	//创建和原迷宫同样大小的矩阵信息。用于记录
	steps := make([][]int, len(maze))

	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}

	// point 队列。用于记录当前可以走的 point --》即 point 在迷宫中对应的值不为 1 的 point
	Q := []point{start}

	// 只要队列中还有值，就继续进行取第一个值进行'广度优先' 上下左右四点的探索，并把可以走的点加入到队列中
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		//如果已经到终点了就直接退出，结束
		if cur == end {
			break
		}

		//对当前 point 做上下左右四个点的探索，找寻可以继续走的 point  加入队列 Q 中
		for _, dir := range dirs {
			// 根据当前 point 获取其 '上下左右' point 的坐标
			next := cur.add(dir)

			// 判断获取的 next point 是否在迷宫矩阵范围内，
			// 或者该 Point 的坐标位置的值是否是 1 为不可走则跳过
			val, ok := next.isInMazeAndGetPointValue(maze)
			if !ok || val == 1 {
				continue
			}

			// 判断获取的 next point 是否已经在记录已走位置的迷宫矩阵范围内。
			// 或者该 point 的不为 0 说明是已经走过的点了也跳过
			val, ok = next.isInMazeAndGetPointValue(steps)
			if !ok || val != 0 {
				continue
			}

			// 如果是起点的话跳过
			if next == start {
				continue
			}

			// 执行到这里说明，该 cur 对应 "上下左右" 中的某一个位置，也就是该 next point 是可以走的，
			// 将其存入记录迷宫矩阵中标记为探索过可走的点，并记录+1表示是走的第几步
			curSteps, _ := cur.isInMazeAndGetPointValue(steps)
			steps[next.i][next.j] = curSteps + 1

			// 最后加入可走队列中
			Q = append(Q, next)
		}
	}

	return steps

}

func main() {
	// 1、从文件中读取迷宫矩阵信息
	maze := readMaze(currentFile() + "maze.in")

	// 打印下读取出来的迷宫矩阵信息
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	steps := walkMaze(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	fmt.Println("===================")

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
