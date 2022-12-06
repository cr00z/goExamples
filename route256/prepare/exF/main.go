package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getFriends(relations [][5]int) []map[int]int {
	friends := make([]map[int]int, len(relations))
	for userIdx, user := range relations {
		friends[userIdx] = make(map[int]int)
		for _, friend := range user {
			if friend == 0 {
				break
			}
			for _, friend2 := range relations[friend-1] {
				if friend2 == 0 {
					break
				}
				if friend2 == userIdx+1 || friend2 == user[0] || friend2 == user[1] ||
					friend2 == user[2] || friend2 == user[3] || friend2 == user[4] {
					continue
				}
				if _, inMap := friends[userIdx][friend2]; inMap {
					friends[userIdx][friend2]++
				} else {
					friends[userIdx][friend2] = 1
				}
			}
		}
	}
	return friends
}

func main() {
	var users, pairs int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &users, &pairs)
	relations := make([][5]int, users)
	for i := 0; i < pairs; i++ {
		var fr1, fr2 int
		fmt.Fscan(in, &fr1, &fr2)
		for idx := range relations[fr1-1] {
			if relations[fr1-1][idx] == 0 {
				relations[fr1-1][idx] = fr2
				break
			}
		}
		for idx := range relations[fr2-1] {
			if relations[fr2-1][idx] == 0 {
				relations[fr2-1][idx] = fr1
				break
			}
		}
	}
	friends := getFriends(relations)
	//fmt.Println(friends)
	for _, user := range friends {
		max := 0
		for _, val := range user {
			if val > max {
				max = val
			}
		}
		friends := make([]int, 0)
		for key, val := range user {
			if val == max {
				friends = append(friends, key)
			}
		}
		if len(friends) == 0 {
			fmt.Println(0)
		} else {
			sort.Ints(friends)
			fmt.Println(strings.Trim(fmt.Sprint(friends), "[]"))
		}
	}
}
