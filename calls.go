package mockargs

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// Calls is a slice of Args that represents
// multiple calls to a number of functions
type Calls []Args

// Equal defines equality for Calls, using reflect package
// calls Equal for each set of Args
func (c Calls) Equal(o Calls, opts ...cmp.Option) error {
	if c == nil && o == nil {
		return nil
	}
	if len(opts) < 1 {
		v := c
		if c == nil {
			v = o
		}
		opts = defaultArguments(v)
	}
	if diff := cmp.Diff(c, o, opts...); diff != "" {
		return fmt.Errorf(diff)
	}
	return nil
}

// In allows asserting wether a function call (defined as set of Args)
// happened in a stack of Calls between certain indexes
func (c Calls) In(a Args, start, end int) error {
	// TODO: maybe just panic? its for tests, actually cleaner to panic maybe
	if start < 0 || end < 0 || start > len(c) || end > len(c) {
		return fmt.Errorf("indexes must be between 0 and %d", len(c))
	}
	var s string
	for _, arg := range c[start:end] {
		if err := arg.Equal(a); err == nil {
			return nil
		}
		s += fmt.Sprintf("\n%+v", arg)
	}
	return fmt.Errorf("call\n%+v\n\nnot found in list of calls between %d and %d: %s", a, start, end, s)
}
