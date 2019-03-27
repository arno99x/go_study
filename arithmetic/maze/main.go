package main

import (
	"fmt"
	"os"
)

//每次把当前格子加上下面的矩阵 形成四个新的格式
var dirs = [4]Point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p Point) add(r Point) Point {
	return Point{p.i + r.i, p.j + r.j}
}

type Point struct {
	i, j int
}

/**
读取文件中的迷宫，第一行是标文件的行数与列数；后面的矩阵数据 0表示是路，1表示是墙
*/
func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	fmt.Printf("%d %d\n\n", row, col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func (p Point) at(grid [][]int) (int, bool) {
	//判断行、列是否越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end Point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}
	Q := []Point{start}
	for len(Q) > 0 {
		cur := Q[0]
		if cur == end {
			break
		}
		Q = Q[1:]
		for i := 0; i < len(dirs); i++ {
			nextPoint := cur.add(dirs[i])

			if nextPoint == end {
				fmt.Println()
			}
			//下个点不是0则不能走
			//下个点走过了也不能走
			//下个点为起点也不能走

			//判断nextPoint是否越界或撞墙,如果是则放弃nextPoint
			val, ok := nextPoint.at(maze)
			if val == 1 || !ok {
				continue
			}

			//判断nextPoint是否曾被走过
			val, ok = nextPoint.at(steps)
			if val != 0 {
				continue
			}

			//判断是否回到原点
			if nextPoint == start {
				continue
			}

			//标记steps里的网格被第几步走过
			curSteps, _ := cur.at(steps)
			steps[nextPoint.i][nextPoint.j] = curSteps + 1
			if curSteps == 12 {
				fmt.Println()
			}
			//向队列中放入新的可走路径
			Q = append(Q, nextPoint)
		}
	}
	return steps
}

func main() {
	maze := readMaze("arithmetic/maze/maze.in")
	//for _,item := range maze{
	//	fmt.Printf("%v\n",item)
	//}

	steps := walk(maze,
		Point{0, 0},
		Point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}
