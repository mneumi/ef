package print

import (
	"fmt"

	"github.com/mneumi/ef/core/color"
)

func PrintlnPrimary(format string, args ...interface{}) {
	fmt.Println(color.Cyan(format, args...))
}

func PrintlnSuccess(format string, args ...interface{}) {
	fmt.Println(color.Green(format, args...))
}

func PrintlnWarning(format string, args ...interface{}) {
	fmt.Println(color.Yellow(format, args...))
}

func PrintlnError(format string, args ...interface{}) {
	fmt.Println(color.Red(format, args...))
}
