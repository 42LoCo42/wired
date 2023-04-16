package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/aerogo/aero"
)

type MyAPI struct {
	// atomic types
	Bool   bool
	Int    int
	UInt   uint
	Float  float32
	String string
	// sized (u)ints & float64 also supported

	// structured types
	Array  [8]int
	Slice  []int
	Struct struct {
		Inner string
	}

	// functions
	GetSimple   func()
	GetContext  func(aero.Context) error
	PostSimple  func(any)
	PostContext func(aero.Context, any) error
	Fails       func() error
}

func main() {
	api := MyAPI{
		Bool:   false,
		Int:    -1,
		UInt:   1,
		Float:  0.1,
		String: "abc",

		Array: [8]int{8, 7, 6, 5, 4, 3, 2, 1},
		Slice: []int{1, 3, 3, 7},
		Struct: struct{ Inner string }{
			Inner: "inner",
		},

		GetSimple: func() {
			log.Print("GetSimple called")
		},

		GetContext: func(ctx aero.Context) error {
			msg := fmt.Sprint(
				"GetContext called from ",
				ctx.Request().Internal().RemoteAddr,
			)
			log.Print(msg)
			return ctx.Text(msg)
		},

		PostSimple: func(data any) {
			log.Print("PostSimple called with ", data)
		},

		PostContext: func(ctx aero.Context, data any) error {
			msg := fmt.Sprintf(
				"PostContext called from %s with %s",
				ctx.Request().Internal().RemoteAddr,
				data,
			)
			log.Print(msg)
			return ctx.Text(msg)
		},

		Fails: func() error {
			return errors.New("o no")
		},
	}

	log.Fatal("Registration error: ", RunTheWired(&api))
}
