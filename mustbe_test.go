package mustbe_test

import (
	"errors"
	"fmt"

	"github.com/davidmz/mustbe"
)

func ExampleOK() {
	defer mustbe.Catched(func(err error) {
		fmt.Println("Catched", err)
	})

	err := errors.New("sample error")
	mustbe.OK(err)

	fmt.Println("Will not be printed")
	// Output: Catched sample error
}

func ExampleOKVal() {
	defer mustbe.Catched(func(err error) {
		fmt.Println("Catched", err)
	})

	divide := func(x, y int) (int, error) {
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	}

	fmt.Println("4 / 2 =", mustbe.OKVal(divide(4, 2)).(int))
	fmt.Println("4 / 0 =", mustbe.OKVal(divide(4, 0)).(int)) // will not be printed

	// Output: 4 / 2 = 2
	// Catched division by zero
}

func ExampleOKOr() {
	defer mustbe.Catched(func(err error) {
		fmt.Println("Catched", err)
	})

	var (
		err     error
		goodErr = errors.New("good error")
		badErr  = errors.New("bad error")
	)

	err = goodErr
	fmt.Println(mustbe.OKOr(err, goodErr))

	err = badErr
	fmt.Println(mustbe.OKOr(err, goodErr))

	// Output: good error
	// Catched bad error
}

func ExampleCatchedAs() {
	foo := func() (err error) {
		defer mustbe.CatchedAs(&err)

		mustbe.OK(errors.New("sample error"))
		return nil
	}

	err := foo()
	fmt.Println("Returned", err)
	// Output: Returned sample error
}

func ExampleErrorBag() {
	defer func() {
		if pnc := recover(); pnc != nil {
			if errBag, ok := pnc.(mustbe.ErrorBag); ok {
				fmt.Println("Wrapped error:", errBag.Unwrap())
			}
		}
	}()
	mustbe.Thrown(errors.New("sample error"))
	// Output: Wrapped error: sample error
}
