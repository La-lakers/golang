// maze project main.go
package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	var row, col int
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, //up
	{0, -1}, //left
	{1, 0},  // down
	{0, 1},  //right
}

func (p *point) add(n point) point {
	return point{p.i + n.i, p.j + n.j}
}

func (p *point) at(n [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(n) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(n[p.i]) {
		return 0, false
	}
	return n[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	pointQueue := []point{start}
	for len(pointQueue) > 0 {
		current := pointQueue[0]
		pointQueue = pointQueue[1:]
		if current == end {
			break
		}
		for _, dir := range dirs {
			next := current.add(dir)
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			stepCount, _ := current.at(steps)
			steps[next.i][next.j] = stepCount + 1
			pointQueue = append(pointQueue, next)

		}

	}
	return steps

}

func main() {
	maze := readMaze("maze.in")
	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
	fmt.Println("==================")
	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})
	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%3d", steps[i][j])
		}
		fmt.Println()
	}

}
