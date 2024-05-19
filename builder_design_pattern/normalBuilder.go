package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "WOODEN DOOR"

}
func (b *NormalBuilder) setWindowType() {
	b.windowType = "STEEL WINDOW"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 3
}

func (b *NormalBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}

}
