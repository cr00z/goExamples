package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func getTime(scanner *bufio.Scanner) time.Time {
	scanner.Scan()
	tm, _ := time.Parse("15:04:05", scanner.Text())
	return tm
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	timeA := getTime(scanner)
	timeB := getTime(scanner)
	timeC := getTime(scanner)

	time2 := timeC.Unix() - timeA.Unix()
	if time2 <= 0 {
		timeC = timeC.Add(time.Hour * 24)
		time2 = timeC.Unix() - timeA.Unix()
	}

	if time2%2 == 1 {
		time2++
	}
	correction := time2 / 2
	tm := time.Unix(timeB.Unix()+correction, 0).UTC()
	fmt.Println(tm.Format("15:04:05"))
}
