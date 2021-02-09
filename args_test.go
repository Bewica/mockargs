package mockargs

import (
	"fmt"
	"testing"
)

func TestArgsEq(t *testing.T) {
	table := []struct {
		name   string
		a1, a2 Args
		isEq   bool
	}{
		{name: "zero values", isEq: true},
		{name: "basic Eq", a1: Args{123, "abc"}, a2: Args{123, "abc"}, isEq: true},
		{name: "basic !Eq", a1: Args{123, "abc"}, a2: Args{123, "cde"}, isEq: false},
		{name: "different type !Eq", a1: Args{123, 456}, a2: Args{123, "cde"}, isEq: false},
		{name: "different len !Eq", a1: Args{123, "abc"}, a2: Args{123, "abc", ""}, isEq: false},
		{
			name: "deep Eq",
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
			name: "deep !Eq",
			a1:   Args{123, "abc", Args{123, true, map[string]interface{}{"key": 456}}},
			a2:   Args{123, "abc", Args{123, false, map[string]interface{}{"key": 456}}},
			isEq: false,
		},
		{
			name: "different len deep !Eq",
			a1:   Args{123, "abc", Args{456, false, map[string]interface{}{"key": 456}}},
			a2:   Args{123, "abc", Args{456, false, true, map[string]interface{}{"key": 456}}},
			isEq: false,
		},
		{name: "func ignored", a1: Args{123, "abc", fmt.Sprintf}, a2: Args{123, "abc", fmt.Sprintf}, isEq: true},
	}
	for idx, test := range table {
		t.Run(fmt.Sprintf("%d - TestArgsEq: %s", idx, test.name), func(t *testing.T) {
			err := test.a1.Eq(test.a2)
			if test.isEq != (err == nil) {
				t.Fatalf("expected Eq to be %t and got: %v", test.isEq, err)
			}
		})
	}
}