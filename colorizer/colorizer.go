package colorizer

import (
	"fmt"
)

const (
	black   = "\033[5;30m%s\033[0m"
	red     = "\033[5;31m%s\033[0m"
	green   = "\033[5;32m%s\033[0m"
	yellow  = "\033[5;33m%s\033[0m"
	purple  = "\033[5;34m%s\033[0m"
	magenta = "\033[5;35m%s\033[0m"
	teal    = "\033[5;36m%s\033[0m"
	white   = "\033[1;37m%s\033[0m"
)

// Black is
func Black(args ...interface{}) string {
	return fmt.Sprintf(black, fmt.Sprint(args...))
}

// Red is
func Red(args ...interface{}) string {
	return fmt.Sprintf(red, fmt.Sprint(args...))
}

// Green is
func Green(args ...interface{}) string {
	return fmt.Sprintf(green, fmt.Sprint(args...))
}

// Yellow is
func Yellow(args ...interface{}) string {
	return fmt.Sprintf(yellow, fmt.Sprint(args...))
}

// Purple is
func Purple(args ...interface{}) string {
	return fmt.Sprintf(red, fmt.Sprint(args...))
}

// Magenta is
func Magenta(args ...interface{}) string {
	return fmt.Sprintf(magenta, fmt.Sprint(args...))
}

// Teal is
func Teal(args ...interface{}) string {
	return fmt.Sprintf(teal, fmt.Sprint(args...))
}

// White is
func White(args ...interface{}) string {
	return fmt.Sprintf(white, fmt.Sprint(args...))
}

// Info is
func Info(name string, args ...interface{}) {
	fmt.Println(fmt.Sprintf("\033[3;36m%s\033[0m[%s]", "INFO", White(name)), Teal(args...))
}

// Warn is
func Warn(name string, args ...interface{}) {
	fmt.Println(fmt.Sprintf("\033[3;33m%s\033[0m[%s]", "WARN", White(name)), Yellow(args...))
}

// Error is
func Error(name string, args ...interface{}) {
	fmt.Println(fmt.Sprintf("\033[3;31m%s\033[0m[%s]", "ERROR", White(name)), Red(args...))
}
