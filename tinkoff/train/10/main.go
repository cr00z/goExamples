package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func squareTriangle(v1, v2, v3 [2]float64) float64 {
	s2 := (v1[0]-v3[0]) * (v2[1]-v3[1]) - (v2[0]-v3[0]) * (v1[1]-v3[1])
	if s2 < 0 {
		s2 = -s2
	}
	return s2 / 2
}

func square(v[][2]float64) float64 {
	var sq float64
	for i := 0; i < len(v)-2; i++ {
		sq += squareTriangle(v[0], v[i+1], v[i+2])
	}
	return sq
}

func main() {
	var n, imin int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	var xmin, xmax float64
	vTmp := make([][2]float64, n, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &vTmp[i][0], &vTmp[i][1])
		if i == 0 {
			xmin, xmax = vTmp[0][0], vTmp[0][0]
			imin = i
		} else {
			if vTmp[i][0] < xmin {
				xmin = vTmp[i][0]
				imin = i
			}
			xmax = math.Max(xmax, vTmp[i][0])
		}
	}

	v := vTmp[imin:]
	v = append(v, vTmp[:imin]...)
	v = append(v, v[0])
	
	//for j:=0;j< 3;j++ {
	for {
		v1 := make([][2]float64, 0)
		v2 := make([][2]float64, 0)
		xpos := xmin + (xmax - xmin)/2
		step := 1
		var y, x, ypos, mtp float64
		for i := 0; i < n + 1; i++ {
			if (v[i][0] < xpos && step == 1) || (step == 3){
				v1 = append(v1, v[i])
			}
			if v[i][0] >= xpos && step == 1 {
				mtp = (xpos - x) / (v[i][0] - x)
				ypos = y + (v[i][1] - y) * mtp
				v1 = append(v1, [2]float64{xpos, ypos})
				v2 = append(v2, [2]float64{xpos, ypos})
				step = 2
			}
			if v[i][0] > xpos && step == 2 {
				v2 = append(v2, v[i])
			}
			if v[i][0] <= xpos && step == 2 {
				mtp = (xpos - x) / (v[i][0] - x)
				ypos = y + (v[i][1] - y) * mtp
				v1 = append(v1, [2]float64{xpos, ypos})
				v2 = append(v2, [2]float64{xpos, ypos})
				step = 3
			}
			y = v[i][1]
			x = v[i][0]
		}
		sq1 := square(v1)
		sq2 := square(v2)
		// fmt.Println(v1, v2)
		if math.Abs(sq1 - sq2) < 0.0000001 {
			fmt.Println(xpos)
			break
		} else {
			if sq1 < sq2 {
				xmin = xpos
			} else {
				xmax = xpos
			}
		}
	}		
}