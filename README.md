# Mockargs

Mockargs is a package that provides convenience functions for testing.
It enables asserting equality around the arguments passed into any Mocked functions
in your tests.

Relies on reflect (and more specifically reflect.DeepEqual) but provides some convenience
wrappers around it and also ignores Func types if they're both != nil.

## TODOs

- [ ] Use [go-cmp](https://github.com/google/go-cmp) as suggested [here](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
