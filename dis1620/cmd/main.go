package main

type Computer struct {
	Core []byte
}

func newComputer() *Computer {
	c := Computer{}
	c.Core = make([]byte, 20000)
}

func (c *Computer) loadCards(fileName string) {

}

func main() {
	c := newComputer()
	c.loadCards("cards.crd")
}