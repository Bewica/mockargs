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

func defaultArguments(v interface{}) cmp.Options {
	opts := cmp.Options{
		cmpopts.EquateEmpty(),
		cmpopts.EquateApprox(5e-4, 0.005),
		cmpopts.EquateApproxTime(time.Second),
		cmpopts.IgnoreTypes(),
	}
	unexporteds := make(map[string]cmp.Option)
	unexporteds = ignoreUnexporteds(unexporteds, v)
	for _, opt := range unexporteds {
		opts = append(opts, opt)
	}
	return opts
}

func ignoreUnexporteds(uniqueIgnores map[string]cmp.Option, v interface{}) map[string]cmp.Option {
	t := reflect.TypeOf(v)
	kind := t.Kind()
	if kind == reflect.Struct {
		uniqueIgnores[t.PkgPath()+t.Name()] = cmpopts.IgnoreUnexported(v)
		return uniqueIgnores
	}
	if kind == reflect.Func {
		uniqueIgnores[t.PkgPath()+t.Name()] = cmpopts.IgnoreTypes(v)
		return uniqueIgnores
	}
	if kind != reflect.Slice {
		return uniqueIgnores
	}
	av := reflect.ValueOf(v)
	l := av.Len()
	for i := 0; i < l; i++ {
		iv := av.Index(i).Interface()
		uniqueIgnores = ignoreUnexporteds(uniqueIgnores, iv)
	}
	return uniqueIgnores
}

// Equal defines equality for Args, using reflect package
// eventually boils down to reflect.DeepEqual but
// ignoring reflect.Func type
func (a Args) Equal(o Args, opts ...cmp.Option) error {
	if a == nil && o == nil {
		return nil
	}
	if len(opts) < 1 {
		opts = defaultArguments(a)
	}
	if diff := cmp.Diff(a, o, opts...); diff != "" {
		return fmt.Errorf(diff)
	}
	return nil
}
