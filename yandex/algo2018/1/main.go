package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var b, c, r, d int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &b, &c, &r, &d)
	bNum := r / 1_000_000
	cNum := r - bNum * 1_000_000
	var kol int64
	for {
		var buy int
		if bNum == 0 {
			buy = c / cNum
		} else if cNum == 0 {
			buy = b / bNum
		} else {
			buy = min(b / bNum, c / cNum)
		}
		kol += int64(buy)
		b -= buy * bNum
		c -= buy * cNum
		d += buy * cNum

		var buy2 int
		if (b >= bNum+1) {
			load := (bNum + 1) * 1_000_000
			sd := load - r
			sdNum := d / sd
			b2Num := b / (bNum + 1)
			buy2 = min(sdNum, b2Num)
			if buy2 > 0 {
				kol += int64(buy2)
				b -= buy2 * (bNum + 1)
				c += buy2 * sd
				d -= buy2 * sd
			}
		}

		if buy == 0 && buy2 == 0 {
			break
		}
	}
	fmt.Println(kol)
}