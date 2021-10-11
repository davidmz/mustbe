// Package mustbe simplifies error handling.
// mustbe.OK* functions receives error argument and panics if is is not nil.
// mustbe.Catched function handle these (and only these) panics.
package mustbe

import (
	"errors"
	"fmt"
)

// ErrorBag is a wrapper around the error value. All OK*/Thrown functions
// are panics with the ErrorBag value. This type is useful for manual panic
// recovering.
type ErrorBag struct{ error }

// Unwrap returns a error wrapped by ErrorBag.
func (e ErrorBag) Unwrap() error { return e.error }

// OK throws panic if err != nil
func OK(err error) {
	if err != nil {
		panic(ErrorBag{err})
	}
}

// Thrown is the synonym of OK
func Thrown(err error) { OK(err) }

// OKVal throws panic if err != nil, oterwise returns val
func OKVal(val interface{}, err error) interface{} {
	if err != nil {
		panic(ErrorBag{err})
	}
	return val
}

// OKOr throws panic if err not nil and not any of errs, oterwise returns err.
// Deprecated: It is recommended to use mustbe.OKOrIs instead.
func OKOr(err error, errs ...error) error {
	if err == nil {
		return nil
	}
	for _, e := range errs {
		if e == err {
			return err
		}
	}
	panic(ErrorBag{err})
}

// OKOrIs throws panic if err not nil and not "Is" (as in errors.Is) any of errs, oterwise returns err
func OKOrIs(err error, errs ...error) error {
	if err == nil {
		return nil
	}
	for _, e := range errs {
		if errors.Is(err, e) {
			return err
		}
	}
	panic(ErrorBag{err})
}

// Catched is a function for defer'ed use (see OK example).
// Errors thrown by mustbe.OK* functions passes to cfun, other panic's are re-panics.
func Catched(cfun func(error)) {
	if pnc := recover(); pnc == nil {
		// none
	} else if eb, ok := pnc.(ErrorBag); ok {
		cfun(eb.error)
	} else {
		panic(pnc)
	}
}

// CatchedAs catches mustbe.* error and assigns it to the targetError.
// It is useful if we just need to return catched error from function.
func CatchedAs(targetError *error) {
	if pnc := recover(); pnc == nil {
		// none
	} else if eb, ok := pnc.(ErrorBag); ok {
		*targetError = eb.error
	} else {
		panic(pnc)
	}
}

// CatchedAs catches mustbe.* error, wraps it using the `fmt.Errorf(format,
// err)` and assigns the result to the targetError. The format string must
// contain a `%w` placeholder for proper fmt.Errorf work.
func CatchedAsAnnotated(targetError *error, format string) {
	if pnc := recover(); pnc == nil {
		// none
	} else if eb, ok := pnc.(ErrorBag); ok {
		*targetError = fmt.Errorf(format, eb.error)
	} else {
		panic(pnc)
	}
}

// True throws panic if test is not true.
func True(test bool, err error) {
	if !test {
		Thrown(err)
	}
}
