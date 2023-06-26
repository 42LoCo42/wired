package main

import (
	"log"
	"reflect"
)

func GetValue(val reflect.Value) any {
	if val.CanUint() {
		res := val.Uint()

		switch val.Type().Size() {
		case 1:
			return uint8(res)
		case 2:
			return uint16(res)
		case 4:
			return uint32(res)
		case 8:
			return uint64(res)
		}
	} else if val.CanInt() {
		res := val.Int()

		switch val.Type().Size() {
		case 1:
			return int8(res)
		case 2:
			return int16(res)
		case 4:
			return int32(res)
		case 8:
			return int64(res)
		}
	} else if val.CanFloat() {
		res := val.Float()

		switch val.Type().Size() {
		case 4:
			return float32(res)
		case 8:
			return float64(res)
		}
	} else if val.CanConvert(reflect.TypeOf("")) {
		return val.String()
	}

	knd := val.Kind()
	switch knd {
	case reflect.Slice:
		return val.Slice(0, val.Len())

	}

	log.New(log.Writer(), "[WARN] ", 0).Printf(
		"GetValue: unknown type: %s (%s)",
		val.Type(),
		val.Kind(),
	)
	return nil
}
