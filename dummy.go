package main

type DummyDrone struct {
	SDKMode bool
	Airborne bool
	VideoOn bool

	Position Position
	Speed int16
}

func NewDrone() *DummyDrone {
	return &DummyDrone{
		SDKMode: false,
		Airborne: false,
		VideoOn: false,

		Position: Position{},
		Speed: 0,
	}
}

type Position struct {
	X float64
	Y float64
	Z float64
	Rotation float64
}