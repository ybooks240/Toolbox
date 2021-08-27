package bench

type MyCat interface {
	Sleep()
	eat()
}

type cat struct {
	name string
	age  int
}

func (c cat) Sleep() {
	panic("not implemented") // TODO: Implement
}

func (c cat) eat() {
	panic("not implemented") // TODO: Implement
}

func test() {

}
