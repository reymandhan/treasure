package main

import "fmt"

// initial map
var treasureMap = [6][8]string{
	{"#", "#", "#", "#", "#", "#", "#", "#"},
	{"#", ".", ".", ".", ".", ".", ".", "#"},
	{"#", ".", "#", "#", "#", ".", ".", "#"},
	{"#", ".", ".", ".", "#", ".", "#", "#"},
	{"#", "X", "#", ".", ".", ".", ".", "#"},
	{"#", "#", "#", "#", "#", "#", "#", "#"},
}

// start position
const x, y int = 1, 4

func init() {
	fmt.Println("Initial Map")
	printMap(0, 0)
}

func main() {
	fmt.Println("\nSearching for Treasure...")
	total := 0
	for i := 1; i <= y; i++ {
		for j := 1; j < len(treasureMap[0])-x; j++ {
			for k := 1; k < len(treasureMap); k++ {
				if y-i+k >= len(treasureMap) {
					break
				}

				newX, newY, status := move(i, j, k)

				if status {
					fmt.Printf("\nPossible treasure location found in (%v,%v) by move north %v step(s), east %v step(s), south %v step(s)\n", newX, newY, i, j, k)
					printMap(newX, newY)
					total++
				}
			}
		}
	}
	fmt.Printf("\nThere are total %v possible treasure location\n", total)
	fmt.Println("Press any key to exit")
	fmt.Scanf("h")
}

// if found clear path, return true and posible treasure position
func move(north, east, south int) (int, int, bool) {
	currPosX, currPosY := x, y

	//moving north
	for i := 1; i <= north; i++ {
		if treasureMap[currPosY-i][currPosX] == "#" {
			//blocked by obstacle
			return -1, -1, false
		}
	}
	currPosY -= north

	//moving east
	for i := 1; i <= east; i++ {
		if treasureMap[currPosY][currPosX+i] == "#" {
			//blocked by obstacle
			return -1, -1, false
		}
	}
	currPosX += east

	//moving south
	for i := 1; i <= south; i++ {
		if treasureMap[currPosY+i][currPosX] == "#" {
			//blocked by obstacle
			return -1, -1, false
		}
	}
	currPosY += south

	return currPosX, currPosY, true
}

func printMap(x, y int) {
	newMap := treasureMap
	if x != 0 && y != 0 {
		newMap[y][x] = "$"
	}

	for _, row := range newMap {
		for _, val := range row {
			fmt.Printf(" %v ", val)
		}
		fmt.Println("")
	}
}
