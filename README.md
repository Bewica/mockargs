# Mockargs

Mockargs is a package that provides convenience functions for testing.
It enables asserting equality around the arguments passed into any Mocked functions
in your tests.

Relies on reflect (and more specifically reflect.DeepEqual) but provides some convenience
wrappers around it and also ignores Func types if they're both != nil.
