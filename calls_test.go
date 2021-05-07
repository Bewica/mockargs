package mockargs

import (
	"fmt"
	"testing"
)

func TestCallsEq(t *testing.T) {
	table := []struct {
		name   string
		c1, c2 Calls
		isEq   bool
	}{
		{name: "zero values", isEq: true},
		{name: "basic Equal", c1: Calls{{123, "abc"}}, c2: Calls{{123, "abc"}}, isEq: true},
		{name: "basic !Equal", c1: Calls{{123, "abc"}}, c2: Calls{{123, "cde"}}, isEq: false},
		{name: "different type !Equal", c1: Calls{{123, "abc"}, {456, 789}}, c2: Calls{{123, "cde"}, {456, "cde"}}, isEq: false},
		{name: "different len !Equal", c1: Calls{{123, "abc"}, {456, 789}}, c2: Calls{{123, "abc"}}, isEq: false},
		{name: "1st calls nil", c1: nil, c2: Calls{{123, "abc"}}, isEq: false},
		{name: "2nd calls nil", c1: Calls{{123, "abc"}}, c2: nil, isEq: false},
	}
	for idx, test := range table {
		t.Run(fmt.Sprintf("%d - TestCallsEq: %s", idx, test.name), func(t *testing.T) {
			err := test.c1.Equal(test.c2)
			if test.isEq != (err == nil) {
				t.Fatalf("expected Equal to be %t and got: %v", test.isEq, err)
			}
		})
	}
}

func TestCallsIn(t *testing.T) {
	table := []struct {
		name  string
		calls Calls
		idxs  [2]int
		call  Args
		isIn  bool
	}{
		{name: "zero values"},
		{
			name:  "basic In",
			calls: Calls{{123, "abc"}, {456, "def"}, {789, "ghi"}},
			idxs:  [2]int{0, 3},
			call:  Args{456, "def"},
			isIn:  true,
		},
		{
			name:  "basic !In",
			calls: Calls{{123, "abc"}, {456, "def"}, {789, "ghi"}},
			idxs:  [2]int{0, 3},
			call:  Args{123, "ghi"},
			isIn:  false,
		},
		{
			name:  "not within indexes !In",
			calls: Calls{{123, "abc"}, {456, "def"}, {789, "ghi"}},
			idxs:  [2]int{0, 2},
			call:  Args{789, "ghi"},
			isIn:  false,
		},
		{
			name:  "wrong indexes",
			calls: Calls{{123, "abc"}, {456, "def"}, {789, "ghi"}},
			idxs:  [2]int{0, 123},
			call:  Args{789, "ghi"},
			isIn:  false,
		},
	}
	for idx, test := range table {
		t.Run(fmt.Sprintf("%d - TestCallsIn: %s", idx, test.name), func(t *testing.T) {
			err := test.calls.In(test.call, test.idxs[0], test.idxs[1])
			if test.isIn != (err == nil) {
				t.Fatalf("expected Equal to be %t and got: %v", test.isIn, err)
			}
		})
	}
}
