package mockargs

import (
	"fmt"
	"testing"
	"time"
)

func TestArgsEq(t *testing.T) {
	table := []struct {
		name   string
		a1, a2 Args
		isEq   bool
	}{
		{name: "zero values", isEq: true},
		{name: "basic Equal", a1: Args{123, "abc"}, a2: Args{123, "abc"}, isEq: true},
		{name: "basic !Equal", a1: Args{123, "abc"}, a2: Args{123, "cde"}, isEq: false},
		{name: "different type !Equal", a1: Args{123, 456}, a2: Args{123, "cde"}, isEq: false},
		{name: "different len !Equal", a1: Args{123, "abc"}, a2: Args{123, "abc", ""}, isEq: false},
		{
			name: "deep Equal",
			a1: Args{123, "abc", Args{struct {
				int
				bool
			}{123, false}, map[string]interface{}{"key": 456}}},
			a2: Args{123, "abc", Args{struct {
				int
				bool
			}{123, false}, map[string]interface{}{"key": 456}}},
			isEq: true,
		},
		{
			name: "deep !Equal",
			a1:   Args{123, "abc", Args{123, true, map[string]interface{}{"key": 456}}},
			a2:   Args{123, "abc", Args{123, false, map[string]interface{}{"key": 456}}},
			isEq: false,
		},
		{
			name: "different len deep !Equal",
			a1:   Args{123, "abc", Args{456, false, map[string]interface{}{"key": 456}}},
			a2:   Args{123, "abc", Args{456, false, true, map[string]interface{}{"key": 456}}},
			isEq: false,
		},
		{name: "func ignored", a1: Args{123, "abc", fmt.Sprintf}, a2: Args{123, "abc", fmt.Sprintf}, isEq: true},
		{
			name: "func if not nil",
			a1:   Args{123, "abc", struct{ fn func() }{fn: func() {}}.fn},
			a2:   Args{123, "abc", struct{ fn func() }{nil}.fn},
			isEq: true,
		},
		{
			name: "deep Equal with Time (1 sec margin)",
			a1:   Args{123, "abc", Args{123, map[string]interface{}{"key": time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC)}}},
			a2:   Args{123, "abc", Args{123, map[string]interface{}{"key": time.Date(1, 2, 3, 4, 5, 6, 700, time.UTC)}}},
			isEq: true,
		},
		{
			name: "deep !Equal with Time (1 sec margin)",
			a1:   Args{123, "abc", Args{123, map[string]interface{}{"key": time.Date(1, 2, 3, 4, 5, 6, 7, time.UTC)}}},
			a2:   Args{123, "abc", Args{123, map[string]interface{}{"key": time.Date(1, 2, 3, 4, 5, 11, 8, time.UTC)}}},
			isEq: false,
		},
		{name: "floats Equal", a1: Args{1.23}, a2: Args{1.234}, isEq: true},
		{name: "floats !Equal", a1: Args{1.23}, a2: Args{1.25}, isEq: false},
		{name: "floats Equal fraction", a1: Args{123.0}, a2: Args{123.05}, isEq: true},
		{name: "floats !Equal fraction", a1: Args{123.0}, a2: Args{123.1}, isEq: false},
		{
			name: "maps Equal",
			a1:   Args{map[string]interface{}{"abc": 12, "cde": 34, "fgh": 56}},
			a2:   Args{map[string]interface{}{"abc": 12, "cde": 34, "fgh": 56}},
			isEq: true,
		},
		{
			name: "maps !Equal",
			a1:   Args{map[string]interface{}{"abc": 12, "cde": 34, "fgh": 56}},
			a2:   Args{map[string]interface{}{"abc": 12, "cde": 43, "fgh": 56}},
			isEq: false,
		},
		{name: "1st arg is nil", a1: nil, a2: Args{123, "abc"}, isEq: false},
		{name: "2nd arg is nil", a1: Args{123, "abc"}, a2: nil, isEq: false},
		{name: "nil in slices", a1: Args{[]interface{}{1, 2, nil}}, a2: Args{123, "abc"}, isEq: false},
	}
	for idx, test := range table {
		t.Run(fmt.Sprintf("%d - TestArgsEq: %s", idx, test.name), func(t *testing.T) {
			err := test.a1.Equal(test.a2)
			if test.isEq != (err == nil) {
				t.Fatalf("expected Equal to be %t and got: %v", test.isEq, err)
			}
		})
	}
}
