package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var n int16 = 41

func run(curr Cords, grid [][]string, stack Stack) {
	grid[curr.y][curr.x] = " "
	// print_grid(grid)

	var next Cords
	var found bool
	next, found = get_neighbours(curr, grid)

	if found {
		stack.Push(curr)

		if curr.y == next.y {
			if next.x > curr.x {
				grid[curr.y][curr.x+1] = " "
			} else {
				grid[curr.y][curr.x-1] = " "
			}
		} else {
			if next.y > curr.y {
				grid[curr.y+1][curr.x] = " "
			} else {
				grid[curr.y-1][curr.x] = " "
			}
		}
		run(next, grid, stack)
	} else if !stack.IsEmpty() {
		run(stack.Pop(), grid, stack)
	}
}

func get_args() (int16, bool) {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main.exe <maze_size>")
		return 0, true
	}

	na := os.Args[1]
	t, err := strconv.Atoi(na)

	if err != nil {
		fmt.Println("Invalid maze size:", na)
		return 0, true
	}

	return int16(t), false
}

func main() {
	startt := time.Now()

	// get the value of n
	tmp, corr := get_args()
	if corr {
		return
	}
	n = tmp

	var stack Stack
	grid := create_grid(n)

	var start Cords
	start.x = 1
	start.y = 1

	run(start, grid, stack)
	elapsed := time.Since(startt)

	print_grid(grid)

	err := save_maze(grid)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Maze saved successfully.")
	}

	fmt.Printf("Program took %s to complete.\n", elapsed)

}
