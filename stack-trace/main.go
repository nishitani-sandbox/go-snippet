package main

type Exampler interface {
	Example([]string, string, int)
}

type E struct{}

func (e *E) Example(s []string, str string, i int) {
	panic("Want error")
}

func main() {
	var e E
	example(&e)
}

func example(e Exampler) {
	s := make([]string, 2, 4)
	e.Example(s, "hello", 10)
}
