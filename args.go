package mockargs

import (
	"fmt"
	"reflect"
)

// Args is a slice of interface{} that represents
// any number of args passed into a function
type Args []interface{}

// Eq defines equality for Args, using reflect package
// eventually boils down to reflect.DeepEqual but
// ignoring reflect.Func type
func (a Args) Eq(o Args) error {
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
			// ignore functions
			continue
		}
		if t.Kind() != reflect.Slice {
			// time comparisons are killing this
			if !reflect.DeepEqual(arg, oarg) {
				return fmt.Errorf("different argument %d, got:\n%+v\nand:\n%+v", adx, arg, oarg)
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
		if err := Args(as).Eq(Args(os)); err != nil {
			return fmt.Errorf("different argument %d, got:\n%+v\nand:\n%+v\nfrom:\n%w", adx, arg, oarg, err)
		}
	}
	return nil
}
