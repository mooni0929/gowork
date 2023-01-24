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
func InterfaceCallAdd() {
	println("Call interface 최신버전 반영 테스트@")
}
func InterfaceCallAddV2() {
	println("Call interface 최신버전 반영 테스트v2")
}
func InterfaceCallAddV6() {
	println("Call interface 최신버전 반영 테스트v6")
}
func InterfaceCallAddV7() {
	println("Call interface 최신버전 반영 테스트v7")
}
func InterfaceCallAddV8() {
	println("Call interface 최신버전 반영 테스트v8")
}
