package alfred

import (
	"fmt"
	"io"
	"os"
)

var Device = os.Stderr

func Log(format string, value ...interface{}) {
	format += "\n"
	log(Device, format, value...)
}

func log(writer io.Writer, format string, value ...interface{}) {
	fmt.Fprintf(writer, format, value...)
}

func Deliver(str string) {
	fmt.Println(str)
}
