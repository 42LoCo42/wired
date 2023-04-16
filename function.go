package main

import (
	"net/http"
	"reflect"

	"github.com/aerogo/aero"
	"github.com/pkg/errors"
)

var errInvalidFunction = errors.New("invalid function type")

func GetFunctionInputs(function reflect.Type) (
	useContext, useData bool, err error) {
	numIn := function.NumIn()

	switch numIn {
	case 0:
		return false, false, nil

	case 1:
		first := function.In(0)

		if first == ctxTyp {
			return true, false, nil
		} else if first == anyTyp {
			return false, true, nil
		} else {
			return false, false, errInvalidFunction
		}

	case 2:
		first := function.In(0)
		second := function.In(1)

		if first == ctxTyp && second == anyTyp {
			return true, true, nil
		} else {
			return false, false, errInvalidFunction
		}
	}

	return false, false, errInvalidFunction
}

func RegisterFunction(
	app *aero.Application,
	path string,
	function reflect.Value,
) error {
	useContext, useData, err := GetFunctionInputs(function.Type())
	if err != nil {
		return errors.Wrap(err, "invalid inputs")
	}

	method := app.Get
	if useData {
		method = app.Post
	}

	method(path, func(ctx aero.Context) error {
		in := []reflect.Value{}
		if useContext {
			in = append(in, reflect.ValueOf(ctx))
		}
		if useData {
			data, err := ctx.Request().Body().JSON()
			if err != nil {
				return ctx.Error(http.StatusBadRequest, err)
			}

			in = append(in, reflect.ValueOf(data))
		}

		out := function.Call(in)
		if len(out) > 0 {
			if err, _ := out[0].Interface().(error); err != nil {
				return ctx.Error(http.StatusInternalServerError, err)
			}
		}

		return nil
	})

	return nil
}
