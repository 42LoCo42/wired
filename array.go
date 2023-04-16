package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/aerogo/aero"
)

func RegisterArray(
	app *aero.Application,
	path string,
	array reflect.Value,
) {
	app.Get(path+"/:index", func(ctx aero.Context) error {
		index, err := strconv.ParseUint(ctx.Get("index"), 10, 64)
		if err != nil {
			return ctx.Error(http.StatusBadRequest, err)
		}

		if index >= uint64(array.Len()) {
			return ctx.Error(http.StatusBadRequest, fmt.Sprintf(
				"index %d out of range (max %d)",
				index, array.Len()-1,
			))
		}

		return ctx.JSON(FieldPTR(array.Index(int(index))))
	})
}
