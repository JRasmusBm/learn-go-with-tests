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

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	for _, x := range haystack {
		if x == needle {
			return
		}
	}

	t.Errorf("Expected %v to contain %q but it didn't", haystack, needle)
}

func setupChannel(addresses []Address) chan Address {
	result := make(chan Address)

	go func() {
		for _, address := range addresses {
			result <- address
		}
		close(result)
	}()

	return result
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
		{
			Name: "Arrays",
			Input: [2]Address{
				{"Storgatan", 33},
				{"Lillgatan", 11},
			},
			ExpectedCalls: []string{"Storgatan", "Lillgatan"},
		},
		{
			Name: "Channels",
			Input: setupChannel([]Address{
				{"Storgatan", 33},
				{"Lillgatan", 11},
			}),
			ExpectedCalls: []string{"Storgatan", "Lillgatan"},
		},
		{
			Name: "Functions",
			Input: func() (Address, Address) {
				return Address{"Storgatan", 33}, Address{"Lillgatan", 11}
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

	t.Run("Maps", func(t *testing.T) {
		var got []string

		walk(map[string]Address{
			"Peter": {"Storgatan", 33},
			"John":  {"Lillgatan", 11},
		}, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Storgatan")
		assertContains(t, got, "Lillgatan")
	})
}
