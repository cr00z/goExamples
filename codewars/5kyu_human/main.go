package main

import (
	"fmt"
)

func HumanReadableTime(seconds int) string {
	return fmt.Sprintf("%d:%d:%d", seconds/3600, (seconds%3600)/60, (seconds%3600)%60)
}

func main() {
	fmt.Println(HumanReadableTime(45296))
}
