package walk

import (
	"reflect"
	"testing"
)

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
