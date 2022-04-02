package format

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"math/rand"
	"time"
)

func Rand() func(w io.Writer, a ...interface{}) {
	c := int(color.FgBlack) + rand.Intn(8)

	return color.New(color.Attribute(c)).FprintFunc()
}

func NoColor() func(w io.Writer, a ...interface{}) {
	return func(w io.Writer, a ...interface{}) {
		fmt.Fprint(w, a...)
	}
}

var (
	Black   = color.New(color.FgBlack).FprintFunc()
	White   = color.New(color.FgWhite).FprintFunc()
	Green   = color.New(color.FgGreen).FprintFunc()
	Yellow  = color.New(color.FgYellow).FprintFunc()
	Cyan    = color.New(color.FgCyan).FprintFunc()
	Magenta = color.New(color.FgMagenta).FprintFunc()
	Red     = color.New(color.FgRed).FprintFunc()
	Blue    = color.New(color.FgBlue).FprintFunc()
)

func (j *JsonFormatter) initColor() {
	rand.Seed(time.Now().Unix())
	r1 := Rand()
	j.numberColor = func(a ...interface{}) {
		r1(j.w, a...)
	}

	r2 := Rand()
	j.mapKeyColor = func(a ...interface{}) {
		r2(j.w, a...)
	}

	r3 := Rand()
	j.boolColor = func(a ...interface{}) {
		r3(j.w, a...)
	}

	r4 := Rand()
	j.nullColor = func(a ...interface{}) {
		r4(j.w, a...)
	}

	r5 := Rand()
	j.stringColor = func(a ...interface{}) {
		r5(j.w, a...)
	}

	r6 := NoColor()
	j.noColor = func(a ...interface{}) {
		r6(j.w, a...)
	}
}

func (j *JsonFormatter) DisableColor() *JsonFormatter {
	r := NoColor()
	j.numberColor = func(a ...interface{}) {
		r(j.w, a...)
	}

	j.mapKeyColor = func(a ...interface{}) {
		r(j.w, a...)
	}

	j.boolColor = func(a ...interface{}) {
		r(j.w, a...)
	}

	j.nullColor = func(a ...interface{}) {
		r(j.w, a...)
	}

	j.stringColor = func(a ...interface{}) {
		r(j.w, a...)
	}

	return j
}

func (j *JsonFormatter) NumberColor(fn func(w io.Writer, a ...interface{})) *JsonFormatter {
	j.numberColor = func(a ...interface{}) {
		fn(j.w, a...)
	}
	return j
}

func (j *JsonFormatter) MapKeyColor(fn func(w io.Writer, a ...interface{})) *JsonFormatter {
	j.mapKeyColor = func(a ...interface{}) {
		fn(j.w, a...)
	}
	return j
}
func (j *JsonFormatter) BoolColor(fn func(w io.Writer, a ...interface{})) *JsonFormatter {
	j.boolColor = func(a ...interface{}) {
		fn(j.w, a...)
	}
	return j
}
func (j *JsonFormatter) NullColor(fn func(w io.Writer, a ...interface{})) *JsonFormatter {
	j.nullColor = func(a ...interface{}) {
		fn(j.w, a...)
	}
	return j
}
func (j *JsonFormatter) StringColor(fn func(w io.Writer, a ...interface{})) *JsonFormatter {
	j.stringColor = func(a ...interface{}) {
		fn(j.w, a...)
	}
	return j
}
