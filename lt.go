package librariesTest

type LibrariesTest struct {
	A string
	B string
}

func NewLibrariesTest() *LibrariesTest {
	return &LibrariesTest{}
}

func (l LibrariesTest) Hello() {
	println("Hello")
}
