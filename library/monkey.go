package library

type PrintFn func(string) string

func Formatter(input string) string {
	return input + "!"
}

func HelloWorld(input string) string {
	return Formatter(input)
}
