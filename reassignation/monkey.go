package reassignation

type PrintFn func(string) string

var PrinterFn PrintFn = Formatter

func Formatter(input string) string {
	return input + "!"
}

func HelloWorld(input string) string {
	return PrinterFn(input)
}
