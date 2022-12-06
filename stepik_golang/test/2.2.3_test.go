package main

import (
	"testing"
	"fmt"
)

func __TestIntMin(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {1, 1, 1},
    }

    for _, test := range tests {
        name := fmt.Sprintf("case(%d,%d)", test.a, test.b)
        t.Run(name, func(t *testing.T) {
            got := IntMin(test.a, test.b)
            if got != test.want {
                t.Errorf("got %d, want %d", got, test.want)
            }
        })
    }
}