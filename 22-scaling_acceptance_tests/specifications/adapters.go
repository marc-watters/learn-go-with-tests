package specifications

type GreetAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}

type IntroduceAdapter func(name string) string

func (i IntroduceAdapter) Introduce(name string) (string, error) {
	return i(name), nil
}
