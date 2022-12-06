package main

import (
	"fmt"
)

func overlap(in *[][2]int) bool {
	result := false
	for i := 0; i < len(*in); i++ {
		for j := i + 1; j < len(*in); j++ {
			if (((*in)[i][0] <= (*in)[j][0]) && ((*in)[i][1] >= (*in)[j][0])) ||
				(((*in)[j][0] <= (*in)[i][0]) && ((*in)[j][1] >= (*in)[i][0])) {
				if (*in)[j][0] < (*in)[i][0] {
					(*in)[i][0] = (*in)[j][0]
					result = true
				}
				if (*in)[j][1] > (*in)[i][1] {
					(*in)[i][1] = (*in)[j][1]
					result = true
				}
				(*in)[j] = [2]int{0, 0}
			}
		}
	}
	return result
}

func SumOfIntervals(in [][2]int) int {
	for overlap(&in) {
	}
	result := 0
	for i := 0; i < len(in); i++ {
		result += in[i][1] - in[i][0]
	}
	return result
}

func main() {
	fmt.Println(SumOfIntervals([][2]int{{-370, -203}, {439, 498}, {33, 45}, {154, 201}, {-181, 429}, {221, 256}}))
}
