package main

import (
	"bufio"
	"fmt"
	"os"
	"container/list"
)

type photoPart struct {
	minx, miny, maxx, maxy int
}

type photo struct {
	minx, miny, maxx, maxy int
	parts *list.List
}

func intersectPhoto(photo photoPart, minx, miny, maxx, maxy int) (*[3]photoPart, *photoPart) {
	return nil, nil
}

func intersectPhotos(photos []*photo, count, minx, miny, maxx, maxy int) {
	for i := 0; i < count; i++ {
		// если не пересекаются
		if minx > photos[i].maxx || maxx < photos[i].minx ||
			miny > photos[i].maxy || maxy < photos[i].miny {
				continue
			}
		for e := photos[i].parts.Front(); e != nil; e = e.Next() {
			if part, isOk := e.Value.(photoPart); isOk {
				append, part2 := intersectPhoto(part, minx, miny, maxx, maxy)
				photos[i].parts.PushBack(append[0])
				photos[i].parts.PushBack(append[1])
				photos[i].parts.PushBack(append[2])
				if part2 != nil {
					append, _ := intersectPhoto(*part2, minx, miny, maxx, maxy)
				}
			}
		}
	}
}

func countArea(photos []*photo) {
	for _, photo := range photos {
		count := 0
		for e := photo.parts.Front(); e != nil; e = e.Next() {
			if part, isOk := e.Value.(photoPart); isOk {
				count += (part.maxx - part.minx) * (part.maxy - part.miny)
			}
		}
		fmt.Println(count)
	}
}

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	photos := make([]*photo, n, n)
	for i := 0; i < n; i++ {
		var minx, miny, maxx, maxy int
		fmt.Fscan(in, &minx, &miny, &maxx, &maxy)
		photos[i] = &photo{minx, miny, maxx, maxy, list.New()}
		photos[i].parts.PushBack(photoPart{minx, miny, maxx, maxy})
		intersectPhotos(photos, i, minx, miny, maxx, maxy)
	}
	countArea(photos)
}