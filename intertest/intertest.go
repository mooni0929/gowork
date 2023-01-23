package intertest

type SomeStruct struct {
	Name string
}
type Hello interface {
	Run()
}

func (s *SomeStruct) Hi() {
	println("Hi " + s.Name)
}
