package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/aerogo/aero"
)

var (
	// why can't I just use TypeOf(aero.Context), this is so stupid
	ctxTyp = reflect.TypeOf(
		aero.Handler(func(ctx aero.Context) error {
			return nil
		}),
	).In(0)

	anyTyp = reflect.TypeOf(reflect.TypeOf).In(0)
)

func RunTheWired(api any) error {
	app := aero.New()
	if err := RegisterStruct(
		app, "", reflect.ValueOf(api).Elem(),
	); err != nil {
		return err
	}

	app.Run()
	return nil
}

func RegisterStruct(
	app *aero.Application,
	prefix string,
	val reflect.Value,
) error {
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		kind := fieldVal.Kind()
		path := prefix + "/" + val.Type().Field(i).Name

		if kind == reflect.Struct {
			RegisterStruct(app, path, fieldVal)
			continue // don't print struct groups
		}

		log.Print("Registering " + path)

		if kind == reflect.Array || kind == reflect.Slice {
			RegisterArray(app, path, fieldVal)

		} else if kind == reflect.Func {
			if err := RegisterFunction(app, path, fieldVal); err != nil {
				return err
			}

		} else if fieldPtr := FieldPTR(fieldVal); fieldPtr != nil {
			RegisterAtomic(app, path, fieldVal, fieldPtr)

		} else {
			return fmt.Errorf(
				"unknown kind %s in field %s",
				kind, fieldVal,
			)
		}
	}

	return nil
}
