// Package mustbe simplifies error handling.
// mustbe.OK* functions receives error argument and panics if is is not nil.
// mustbe.Catched function handle these (and only these) panics.
package mustbe

type errorBag struct{ error }

// OK throws panic if err != nil
func OK(err error) {
	if err != nil {
		panic(errorBag{err})
	}
}

// OKVal throws panic if err != nil, oterwise returns val
func OKVal(val interface{}, err error) interface{} {
	if err != nil {
		panic(errorBag{err})
	}
	return val
}

// OKOr throws panic if err not nil and not in errs, oterwise returns err
func OKOr(err error, errs ...error) error {
	if err == nil {
		return nil
	}
	for _, e := range errs {
		if e == err {
			return err
		}
	}
	panic(errorBag{err})
}

// Catched is a function for defer'ed use (see OK example).
// Errors thrown by mustbe.OK* functions passes to cfun, other panic's are re-panics.
func Catched(cfun func(error)) {
	if pnc := recover(); pnc == nil {
		// none
	} else if eb, ok := pnc.(errorBag); ok {
		cfun(eb.error)
	} else {
		panic(pnc)
	}
}
