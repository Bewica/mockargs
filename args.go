package mockargs

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Args is a slice of interface{} that represents
// any number of args passed into a function
type Args []interface{}

// Equal defines equality for Args, using reflect package
// eventually boils down to reflect.DeepEqual but
// ignoring reflect.Func type
func (a Args) Equal(o Args, opts ...cmp.Option) error {
	if a == nil && o == nil {
		return nil
	}
	if a == nil {
		return fmt.Errorf("got\n%+v\nand\n%+v", a, o)
	}
	if o == nil {
		return fmt.Errorf("got\n%+v\nand\n%+v", a, o)
	}
	if len(a) != len(o) {
		return fmt.Errorf("got different number of arguments: %d and %d", len(a), len(o))
	}
	for adx, arg := range a {
		oarg := o[adx]
		t := reflect.TypeOf(arg)
		ot := reflect.TypeOf(oarg)
		if t != ot {
			return fmt.Errorf("different argument %d, got:\n%+v\nand:\n%+v", adx, arg, oarg)
		}
		if t.Kind() == reflect.Func {
			argIsNil := reflect.ValueOf(arg).IsNil()
			oargIsNil := reflect.ValueOf(oarg).IsNil()
			if argIsNil != oargIsNil {
				return fmt.Errorf("different argument %d, got:\n%+v\nand:\n%+v", adx, arg, oarg)
			}
			// ignore functions (if both nil or != nil)
			continue
		}
		if t.Kind() != reflect.Slice {
			if opts == nil {
				opts = cmp.Options{cmpopts.EquateApproxTime(5 * time.Second), cmpopts.EquateApprox(5e-4, 0.005)}
			}
			if t.Kind() == reflect.Struct {
				opts = append(opts, cmpopts.IgnoreUnexported(arg))
			}
			if diff := cmp.Diff(arg, oarg, opts...); diff != "" {
				return fmt.Errorf("different argument %d:\n%s", adx, diff)
			}
			continue
		}
		av := reflect.ValueOf(arg)
		ov := reflect.ValueOf(oarg)
		if av.Len() != ov.Len() {
			return fmt.Errorf("different argument %d, got:\n%+v\nand:\n%+v", adx, arg, oarg)
		}
		as := make([]interface{}, av.Len())
		os := make([]interface{}, ov.Len())
		for i := 0; i < av.Len(); i++ {
			as[i] = av.Index(i).Interface()
			os[i] = ov.Index(i).Interface()
		}
		if err := Args(as).Equal(Args(os)); err != nil {
			return fmt.Errorf("different argument %d, got:\n%+v\nand:\n%+v\nfrom:\n%w", adx, arg, oarg, err)
		}
	}
	return nil
}
