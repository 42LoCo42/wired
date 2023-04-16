package main

import (
	"net/http"
	"reflect"

	"github.com/aerogo/aero"
)

func RegisterAtomic(
	app *aero.Application,
	path string,
	field reflect.Value,
	fieldPtr any,
) {
	typ := field.Type()

	app.Get(path, func(ctx aero.Context) error {
		return ctx.JSON(fieldPtr)
	})

	app.Post(path, func(ctx aero.Context) error {
		data, err := ctx.Request().Body().JSON()
		if err != nil {
			return ctx.Error(http.StatusBadRequest, err)
		}

		value := reflect.ValueOf(data)
		if !value.CanConvert(typ) {
			return ctx.Error(http.StatusBadRequest, "value of invalid type")
		}

		field.Set(value.Convert(typ))
		return nil
	})

}

func FieldPTR(value reflect.Value) any {
	if !value.CanAddr() {
		return nil
	}

	ptr := value.Addr().UnsafePointer()
	switch value.Kind() {
	case reflect.Bool:
		return (*bool)(ptr)
	case reflect.Int:
		return (*int)(ptr)
	case reflect.Int8:
		return (*int8)(ptr)
	case reflect.Int16:
		return (*int16)(ptr)
	case reflect.Int32:
		return (*int32)(ptr)
	case reflect.Int64:
		return (*int64)(ptr)
	case reflect.Uint:
		return (*uint)(ptr)
	case reflect.Uint8:
		return (*uint8)(ptr)
	case reflect.Uint16:
		return (*uint16)(ptr)
	case reflect.Uint32:
		return (*uint32)(ptr)
	case reflect.Uint64:
		return (*uint64)(ptr)
	case reflect.Float32:
		return (*float32)(ptr)
	case reflect.Float64:
		return (*float64)(ptr)
	case reflect.String:
		return (*string)(ptr)
	}

	return nil
}
