package main_

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

func main_() {
	var ul, sl, duration, tm, uid int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &ul, &sl, &duration)

	total := make([]int, 0, tmQueueSize)
	tmMap := make(map[int][]int)

	for {
		fmt.Fscan(in, &tm, &uid)
		if tm == -1 {
			break
		}

		userTmQueue, inMap := tmMap[uid]
		if !inMap {
			userTmQueue = []int{tm}
		} else {
			userTmQueue = append(userTmQueue, tm)
			for userTmQueue[0] < tm-duration {
				userTmQueue = userTmQueue[1:]
			}
		}
		tmMap[uid] = userTmQueue

		total = append(total, tm)
		for total[0] < tm-duration {
			total = total[1:]
		}

		if len(userTmQueue) > ul {
			outResponseCode(out, http.StatusTooManyRequests)

			continue
		}

		if len(total) > sl {
			outResponseCode(out, http.StatusServiceUnavailable)
			continue
		}

		outResponseCode(out, http.StatusOK)
	}
}
