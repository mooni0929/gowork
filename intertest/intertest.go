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
func InterfaceCall() {
	println("Call interface 최신버전 반영 테스트")
}
