package reflections

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
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
			}{
				Name: "Anuj",
			},
			ExpectedCalls: []string{"Anuj"},
		},
		{
			Name: "Struct with 2 string fields",
			Input: struct {
				Name string
				City string
			}{
				Name: "Anuj",
				City: "Delhi",
			},
			ExpectedCalls: []string{"Anuj", "Delhi"},
		},
		{
			Name: "Struct with one integer field",
			Input: struct {
				Name string
				Age  int
			}{
				Name: "Anuj",
				Age:  30,
			},
			ExpectedCalls: []string{"Anuj"},
		},
		{
			Name: "Nested fields",
			Input: struct {
				Name    string
				Profile Profile
			}{
				Name: "Anuj",
				Profile: Profile{
					Age:  30,
					City: "Delhi",
				},
			},
			ExpectedCalls: []string{"Anuj", "Delhi"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got %v, wanted %v", got, tt.ExpectedCalls)
			}
		})
	}
}
