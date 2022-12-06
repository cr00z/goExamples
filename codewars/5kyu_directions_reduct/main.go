package main

import (
	"fmt"
)

func DirReduc(arr []string) []string {
	path := make([]string, 0)
	for _, dir := range arr {
		var opposite string = ""
		switch dir {
		case "NORTH":
			opposite = "SOUTH"
		case "SOUTH":
			opposite = "NORTH"
		case "WEST":
			opposite = "EAST"
		case "EAST":
			opposite = "WEST"
		}
		if (len(path) > 0) && (path[len(path)-1] == opposite) {
			path = path[:len(path)-1]
		} else {
			path = append(path, dir)
		}
	}
	return path
}

func main() {
	fmt.Println(DirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}))
}
