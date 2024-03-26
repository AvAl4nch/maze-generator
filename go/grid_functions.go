package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func create_grid(n int16) [][]string {
	t := make([][]string, n)
	for i := range t {
		t[i] = make([]string, n)
		for j := range t[i] {
			if i%2 == 0 || j%2 == 0 {
				t[i][j] = "#"
			} else {
				t[i][j] = "0"
			}
		}
	}
	return t
}

func print_grid(grid [][]string) {
	// cmd := exec.Command("cmd", "/c", "cls")
	// cmd.Stdout = os.Stdout
	// cmd.Run()

	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val + " ")
		}
		fmt.Printf("\n")
	}

}

func get_neighbours(curr Cords, grid [][]string) (Cords, bool) {
	unvisited := []Cords{}
	var found bool = false

	if curr.y-2 > 0 && grid[curr.y-2][curr.x] == "0" {
		unvisited = append(unvisited, Cords{y: curr.y - 2, x: curr.x})
		found = true

	}
	if curr.y+2 < n && grid[curr.y+2][curr.x] == "0" {
		unvisited = append(unvisited, Cords{y: curr.y + 2, x: curr.x})
		found = true

	}
	if curr.x-2 > 0 && grid[curr.y][curr.x-2] == "0" {
		unvisited = append(unvisited, Cords{y: curr.y, x: curr.x - 2})
		found = true

	}
	if curr.x+2 < n && grid[curr.y][curr.x+2] == "0" {
		unvisited = append(unvisited, Cords{y: curr.y, x: curr.x + 2})
		found = true
	}

	if len(unvisited) == 0 {
		return Cords{}, false
	}

	index := rand.Intn(len(unvisited))
	return unvisited[index], found
}

func save_maze(grid [][]string) error {
	file, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, row := range grid {
		for _, val := range row {
			_, err := writer.WriteString(val)
			if err != nil {
				return err
			}
		}
		_, err := writer.WriteString("\n")
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
