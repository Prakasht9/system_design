package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}
func (b *IglooBuilder) setWindowType() {
	b.windowType = "SNOW WINDOW"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "SNOW DOOR"

}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}

}
