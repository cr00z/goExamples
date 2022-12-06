package reflection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two strings fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "nested fields",
			Input: Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			ExpectedCalls: []string{"London", "Reykjavik"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			ExpectedCalls: []string{"London", "Reykjavik"},
		},
		{
			Name: "with function",
			Input: func() (Profile, Profile) {
				return Profile{33, "London"}, Profile{34, "Reykjavik"}
			},
			ExpectedCalls: []string{"London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			assert.Equal(t, test.ExpectedCalls, got)
		})
	}
}

func TestWalkMaps(t *testing.T) {
	want := []string{"London", "Reykjavik"}
	input := map[int]string{
		33: "London",
		34: "Reykjavik",
	}
	var got []string
	walk(input, func(input string) {
		got = append(got, input)
	})

	assert.ElementsMatch(t, want, got)
}

func TestWalkChannels(t *testing.T) {
	channel := make(chan Profile)

	go func() {
		channel <- Profile{33, "London"}
		channel <- Profile{34, "Katowice"}
		close(channel)
	}()

	var got []string
	want := []string{"London", "Katowice"}

	walk(channel, func(input string) {
		got = append(got, input)
	})

	assert.Equal(t, want, got)
}
