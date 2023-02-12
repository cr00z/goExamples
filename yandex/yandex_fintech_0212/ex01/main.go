package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const tmQueueSize = 50_000

func outResponseCode(out *bufio.Writer, code int) {
	out.WriteString(strconv.Itoa(code))
	out.WriteByte('\n')
	out.Flush()
}

func main() {
	var userLimit, serviceLimit, duration, currentTime, userID int

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &userLimit, &serviceLimit, &duration)

	total := make([]int, 0, tmQueueSize)
	tmMap := make(map[int][]int)

	for {
		fmt.Fscan(in, &currentTime, &userID)
		if currentTime == -1 {
			break
		}

		userTmQueue, inMap := tmMap[userID]
		if !inMap {
			userTmQueue = make([]int, 0)
		} else {
			for len(userTmQueue) > 0 && (userTmQueue[0] < currentTime-duration) {
				userTmQueue = userTmQueue[1:]
			}
		}

		for len(total) > 0 && (total[0] < currentTime-duration) {
			total = total[1:]
		}

		if len(userTmQueue) >= userLimit {
			outResponseCode(out, http.StatusTooManyRequests)
			continue
		}

		if len(total) >= serviceLimit {
			outResponseCode(out, http.StatusServiceUnavailable)
			continue
		}

		userTmQueue = append(userTmQueue, currentTime)
		tmMap[userID] = userTmQueue

		total = append(total, currentTime)

		outResponseCode(out, http.StatusOK)
	}
}
