package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Street string
	Number int
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{"Rasmus"},
			ExpectedCalls: []string{"Rasmus"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name    string
				Surname string
			}{"Rasmus", "Bergström"},
			ExpectedCalls: []string{"Rasmus", "Bergström"},
		},
		{
			Name: "Struct with non-string field",
			Input: struct {
				Name string
				Age  int
			}{"Rasmus", 28},
			ExpectedCalls: []string{"Rasmus"},
		},
		{
			Name: "Struct with non-string field",
			Input: Person{
				"Rasmus",
				Address{"Storgatan", 33}},
			ExpectedCalls: []string{"Rasmus", "Storgatan"},
		},
		{
			Name: "Pointers to things",
			Input: &Person{
				"Rasmus",
				Address{"Storgatan", 33}},
			ExpectedCalls: []string{"Rasmus", "Storgatan"},
		},
		{
			Name: "Slices",
			Input: []Address{
				{"Storgatan", 33},
				{"Lillgatan", 11},
			},
			ExpectedCalls: []string{"Storgatan", "Lillgatan"},
		},
		{
			Name: "Arrays",
			Input: [2]Address{
				{"Storgatan", 33},
				{"Lillgatan", 11},
			},
			ExpectedCalls: []string{"Storgatan", "Lillgatan"},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			var got []string

			walk(testCase.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, testCase.ExpectedCalls) {
				t.Errorf("Wrong calls, got %v want %v", got, testCase.ExpectedCalls)
			}
		})
	}
}
